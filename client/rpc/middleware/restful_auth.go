package middleware

import (
	"fmt"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"k8s.io/klog/v2"

	"github.com/Aaazj/mcenter/apps/audit"
	"github.com/Aaazj/mcenter/apps/endpoint"
	"github.com/Aaazj/mcenter/conf"

	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/apps/user"
	"github.com/Aaazj/mcenter/client/rpc"
)

// RestfulServerInterceptor go-restful认证中间件
func RestfulServerInterceptor() restful.FilterFunction {
	return newhttpAuther().GoRestfulAuthFunc
}

// 给服务端提供的RESTful接口的 认证与鉴权中间件
func newhttpAuther() *httpAuther {
	return &httpAuther{
		log:    zap.L().Named("auther.http"),
		client: rpc.C(),

		mode: ACL_MODE,
	}
}

type PermissionMode int

const (
	// PRBAC_MODE 基于策略的权限校验
	PRBAC_MODE PermissionMode = 1
	// ACL_MODE 基于用户类型的权限校验
	ACL_MODE PermissionMode = 2
)

type httpAuther struct {
	log logger.Logger
	// 基于rpc客户端进行封装
	client *rpc.ClientSet
	// 鉴权模式
	mode PermissionMode
}

// 设置权限校验策略
func (a *httpAuther) SetPermissionMode(m PermissionMode) {
	a.mode = m
}

// 是否开启权限的控制, 交给中间件使用方去觉得
func (a *httpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}

	// 请求拦截
	entry := endpoint.NewEntryFromRestRequest(req)

	if entry != nil && entry.AuthEnable {
		// 访问令牌校验
		tk, err := a.CheckAccessToken(req)

		if err != nil {
			if err.Error() == token.ErrTokenExpoired.Error() {

				res.Errcode = 444
				res.Errmsg = err.Error()
				klog.V(4).Info(res)
				if err := resp.WriteAsJson(res); err != nil {
					klog.Error(err)
				}
				return
			}

			if err.Error() == token.ErrOtherPlaceLoggedIn.Error() {

				res.Errcode = 50010
				res.Errmsg = err.Error()
				klog.V(4).Info(res)
				if err := resp.WriteAsJson(res); err != nil {
					klog.Error(err)
				}
				return
			}
			res.Errcode = 401001
			res.Errmsg = err.Error()
			klog.V(4).Info(res)
			if err := resp.WriteAsJson(res); err != nil {
				klog.Error(err)
			}
			return
		}

		err = a.CheckPermission(req, tk, entry)
		if err != nil {
			res.Errcode = 401002
			res.Errmsg = err.Error()
			klog.V(4).Info(res)
			if err := resp.WriteAsJson(res); err != nil {
				klog.Error(err)
			}
			return
		}

	}

	start := time.Now()

	// next flow
	next.ProcessFilter(req, resp)

	cost := time.Now().Sub(start).Milliseconds()
	//开启审计
	//Metadata(label.Audit, label.Enable)
	if entry != nil && entry.AuthEnable && entry.AuditLog {

		tk := req.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)
		auditReq := audit.NewOperateLog(tk.Username, "", "")

		auditReq.Url = req.Request.Method + " " + req.Request.URL.String()
		auditReq.Cost = cost
		auditReq.StatusCode = int64(resp.StatusCode())
		auditReq.UserAgent = req.Request.UserAgent()
		// X-Forwar-For
		auditReq.RemoteIp = request.GetRemoteIP(req.Request)

		meta := req.SelectedRoute().Metadata()

		if meta != nil {
			if v, ok := meta[label.Resource]; ok {
				auditReq.Resource, _ = v.(string)
			}
			if v, ok := meta[label.Action]; ok {
				auditReq.Action, _ = v.(string)
			}
		}

		_, err := a.client.Audit().AuditOperate(req.Request.Context(), auditReq)
		if err != nil {
			a.log.Warnf("audit operate failed, %s", err)
			return
		}

	}
}

func (a *httpAuther) CheckAccessToken(req *restful.Request) (*token.Token, error) {
	// 获取用户Token, Token放在Heander Authorization
	ak := token.GetAccessTokenFromHTTP(req.Request)

	if ak == "" {
		ak = token.GetAccessTokenFromSocket(req)
	}

	if ak == "" {
		return nil, token.ErrUnauthorized
	}

	// 调用GRPC 校验用户Token合法性
	tk, err := a.client.Token().ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
	if err != nil {
		return nil, err
	}

	// 是不是需要返回用户的认证信息: 那个人, 那个空间下面， token本身的信息
	req.SetAttribute(token.TOKEN_ATTRIBUTE_NAME, tk)

	return tk, nil
}

func (a *httpAuther) CheckPermission(req *restful.Request, tk *token.Token, e *endpoint.Entry) error {
	if tk == nil {
		return exception.NewUnauthorized("validate permission need token")
	}

	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(user.TYPE_SUPPER) {
		a.log.Debugf("[%s] supper admin skip permission check!", tk.Username)
		return nil
	}

	switch a.mode {
	case ACL_MODE:
		return a.ValidatePermissionByACL(req, tk, e)
	case PRBAC_MODE:
		return a.ValidatePermissionByPRBAC(req, tk, e)
	default:
		return fmt.Errorf("only support acl and prbac")
	}
}

func (a *httpAuther) ValidatePermissionByACL(req *restful.Request, tk *token.Token, e *endpoint.Entry) error {
	// 检查是否是允许的类型
	if len(e.Allow) > 0 {
		a.log.Debugf("[%s] start check permission to keyauth ...", tk.Username)
		if !e.IsAllow(tk.UserType) {
			return exception.NewPermissionDeny("no permission, allow: %s, but current: %s", e.Allow, tk.UserType)
		}
		a.log.Debugf("[%s] permission check passed", tk.Username)
	}

	return nil
}

func (a *httpAuther) ValidatePermissionByPRBAC(r *restful.Request, tk *token.Token, e *endpoint.Entry) error {
	// ci, err := a.client.ClientInfo(r.Request.Context())
	// if err != nil {
	// 	return err
	// }

	// req := policy.NewCheckPermissionRequest()
	// req.Username = tk.Username
	// req.Namespace = tk.Namespace
	// req.ServiceId = ci.Meta.Id
	// req.Path = e.UniquePath()
	// perm, err := a.client.Policy().CheckPermission(r.Request.Context(), req)
	// if err != nil {
	// 	return exception.NewPermissionDeny(err.Error())
	// }
	// a.log.Debugf("[%s] permission check passed", tk.Username)

	// // 保存访问访问信息
	// r.SetAttribute(policy.SCOPE_ATTRIBUTE_NAME, perm.Scope)
	return nil
}
