package api

import (
	"fmt"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"k8s.io/klog/v2"

	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/apps/user"
	"github.com/Aaazj/mcenter/conf"
)

// 主账号用户管理接口

type sub struct {
	service user.Service
	log     logger.Logger
}

func (h *sub) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.service = app.GetInternalApp(user.AppName).(user.Service)
	return nil
}

func (h *sub) Name() string {
	return "account"
}

func (h *sub) Version() string {
	return "v1"
}

func (h *sub) Registry(ws *restful.WebService) {
	tags := []string{"账号管理"}

	ws.Route(ws.POST("/password").To(h.UpdatePassword).
		Metadata(label.Auth, true).
		//Metadata(label.Allow, user.TYPE_SUB).
		Doc("账号修改自己密码").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.UpdatePasswordRequest{}).
		Returns(0, "OK", &user.User{}))
}

func (h *sub) UpdatePassword(r *restful.Request, w *restful.Response) {
	res := conf.GeneralResponse{
		Errcode: 0,
		Errmsg:  "OK",
	}
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)

	req := user.NewUpdatePasswordRequest()
	req.UserId = tk.UserId
	fmt.Printf("req.UserId: %v\n", req.UserId)
	if err := r.ReadEntity(req); err != nil {
		//response.Failed(w, err)
		res.Errmsg = "读取密码信息失败:" + err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}

	set, err := h.service.UpdatePassword(r.Request.Context(), req)
	if err != nil {
		//response.Failed(w, err)
		res.Errcode = 401002
		res.Errmsg = err.Error()
		klog.V(4).Info(res)
		if err := w.WriteAsJson(res); err != nil {
			klog.Error(err)
		}
		return
	}
	res.Data = set
	if err = w.WriteAsJson(res); err != nil {
		klog.Error(err)
	}
	//response.Success(w, set)
}

func init() {
	app.RegistryRESTfulApp(&sub{})
}
