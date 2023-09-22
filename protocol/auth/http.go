package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Aaazj/mcenter/apps/audit"
	"github.com/Aaazj/mcenter/apps/endpoint"
	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/apps/user"
	"github.com/Aaazj/mcenter/conf"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/rs/xid"
	"k8s.io/klog/v2"
)

func NewHttpAuther() *httpAuther {
	return &httpAuther{
		log:              zap.L().Named("auther.http"),
		tk:               app.GetInternalApp(token.AppName).(token.Service),
		audit:            app.GetInternalApp(audit.AppName).(audit.Service),
		cache:            cache.C(),
		codeCheckSilence: 30 * time.Minute,
	}
}

type httpAuther struct {
	log              logger.Logger
	tk               token.Service
	audit            audit.Service
	cache            cache.Cache
	codeCheckSilence time.Duration
}

// 设置静默时长
func (a *httpAuther) SetCodeCheckSilenceTime(t time.Duration) {
	a.codeCheckSilence = t
}

func (a *httpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {

	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}

	// 请求拦截
	entry := endpoint.NewEntryFromRestRequest(req)
	fmt.Printf("\"11111111111111111111111111111111111\": %v\n", "11111111111111111111111111111111111")
	fmt.Printf("entry: %v\n", entry)
	if entry != nil && entry.AuthEnable {
		// 访问令牌校验
		tk, err := a.CheckAccessToken(req)
		if err != nil {
			//response.Failed(resp, err)
			res.Errcode = 401001
			res.Errmsg = err.Error()
			klog.V(4).Info(res)
			if err := resp.WriteAsJson(res); err != nil {
				klog.Error(err)
			}
			return
		}

		// 接口调用权限校验
		err = a.CheckPermission(req.Request.Context(), tk, entry)
		if err != nil {
			response.Failed(resp, err)
			return
		}

	}

	start := time.Now()
	fmt.Printf("start: %v\n", start)
	next.ProcessFilter(req, resp)

	cost := time.Now().Sub(start).Milliseconds()
	//开启审计
	//Metadata(label.Audit, label.Enable)

	if entry != nil && entry.AuthEnable && entry.AuditLog {

		tk := req.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)
		auditReq := audit.NewOperateLog(tk.Username, "", "")

		auditReq.Id = xid.New().String()
		auditReq.Url = req.Request.Method + " " + req.Request.URL.String()
		fmt.Printf("req.Request: %v\n", req.Request)
		fmt.Printf("req.Request.Proto: %v\n", req.Request.Proto)

		fmt.Printf("resp.ResponseWriter: %v\n", resp.ResponseWriter)
		fmt.Printf("resp.ResponseWriter.Header(): %v\n", resp.ResponseWriter.Header())

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
		fmt.Printf("\"aaaaaa\": %v\n", "aaaaaa")
		_, err := a.audit.AuditOperate(req.Request.Context(), auditReq)
		if err != nil {
			a.log.Warnf("audit operate failed, %s", err)
			return
		}

	}
}

func (a *httpAuther) CheckAccessToken(req *restful.Request) (*token.Token, error) {
	// 获取用户Token, Token放在Heander Authorization
	ak := token.GetAccessTokenFromHTTP(req.Request)
	fmt.Printf("aaaaaaak: %v\n", ak)
	if ak == "" {
		return nil, token.ErrUnauthorized
	}
	fmt.Printf("ak: %v\n", ak)
	// 调用GRPC 校验用户Token合法性
	tk, err := a.tk.ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
	if err != nil {
		return nil, err
	}

	// 是不是需要返回用户的认证信息: 那个人, 那个空间下面， token本身的信息
	req.SetAttribute(token.TOKEN_ATTRIBUTE_NAME, tk)
	fmt.Printf("\"111111\": %v\n", "111111")
	return tk, nil
}

func (a *httpAuther) CheckPermission(ctx context.Context, tk *token.Token, e *endpoint.Entry) error {

	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(user.TYPE_SUPPER) {
		a.log.Debugf("[%s] supper admin skip permission check!", tk.Username)
		return nil
	}

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
