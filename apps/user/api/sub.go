package api

import (
	"fmt"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/apps/user"
)

// 主账号用户管理接口

type primary struct {
	service user.Service
	log     logger.Logger
}

func (h *primary) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.service = app.GetInternalApp(user.AppName).(user.Service)
	return nil
}

func (h *primary) Name() string {
	return "user/sub"
}

func (h *primary) Version() string {
	return "v1"
}

func (h *primary) Registry(ws *restful.WebService) {
	tags := []string{"子账号管理"}

	ws.Route(ws.POST("/").To(h.CreateUser).
		Metadata(label.Auth, true).
		Metadata(label.Allow, user.TYPE_PRIMARY).
		Doc("创建子账号").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Returns(200, "创建成功", &user.User{}))

	ws.Route(ws.GET("/").To(h.QueryUser).
		Metadata(label.Auth, true).
		Metadata(label.Allow, user.TYPE_PRIMARY).
		//Metadata(label.Code, true).
		Doc("查询子账号列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", user.UserSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeUser).
		Metadata(label.Auth, true).
		Doc("查询子账号详情").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(user.User{}).
		Returns(200, "OK", user.User{}))

	ws.Route(ws.PUT("/{id}").To(h.PutUser).
		Metadata(label.Auth, true).
		Metadata(label.Allow, user.TYPE_PRIMARY).
		Doc("修改子账号").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchUser).
		Metadata(label.Auth, true).
		Metadata(label.Allow, user.TYPE_PRIMARY).
		Doc("修改子账号").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteUser).
		Metadata(label.Auth, true).
		Metadata(label.Allow, user.TYPE_PRIMARY).
		Doc("删除子账号").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.POST("/{id}/password").To(h.ResetPassword).
		Metadata(label.Auth, true).
		Metadata(label.Allow, user.TYPE_PRIMARY).
		Doc("重置子账号密码").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func (h *primary) CreateUser(r *restful.Request, w *restful.Response) {
	req := user.NewCreateUserRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}
	fmt.Printf("req.Type: %v\n", req.Type)
	fmt.Printf("req.Domain11111: %v\n", req.Domain)
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)
	fmt.Printf("tk: %v\n", tk)
	if tk.UserType != user.TYPE_SUPPER {
		fmt.Printf("int32(req.Type): %v\n", int32(req.Type))
		//没有权限创建超级管理员
		if user.TYPE_name[int32(req.Type)] == "SUPPER" || user.TYPE_name[int32(req.Type)] == "PRIMARY" {
			response.Failed(w, fmt.Errorf("no permission, allow: SUPPER , but current: PRIMARY"))
			return
		}
		req.Domain = tk.Domain
		req.Namespace = tk.Namespace

	} else {
		if req.Domain == "" {
			req.Domain = tk.Domain
		}
		if req.Namespace == "" {
			req.Namespace = tk.Namespace
		}
	}

	fmt.Printf("req.Domain222222222222: %v\n", req.Domain)

	set, err := h.service.CreateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *primary) PutUser(r *restful.Request, w *restful.Response) {
	req := user.NewPutUserRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *primary) PatchUser(r *restful.Request, w *restful.Response) {
	req := user.NewPatchUserRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *primary) ResetPassword(r *restful.Request, w *restful.Response) {
	req := user.NewResetPasswordRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UserId = r.PathParameter("id")

	set, err := h.service.ResetPassword(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *primary) DeleteUser(r *restful.Request, w *restful.Response) {
	req := user.NewDeleteUserRequest()
	req.UserIds = append(req.UserIds, r.PathParameter("id"))

	set, err := h.service.DeleteUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *primary) QueryUser(r *restful.Request, w *restful.Response) {
	req := user.NewQueryUserRequestFromHTTP(r.Request)

	//6.9添加限制
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)
	// fmt.Printf("tk.Domain: %v\n", tk.Domain)
	// fmt.Printf("tk.UserType: %v\n", tk.UserType)
	// fmt.Printf("tk.Namespace: %v\n", tk.Namespace)
	if tk.UserType != user.TYPE_SUPPER {
		req.Domain = tk.Domain
		req.Namespace = tk.Namespace
	}

	ins, err := h.service.QueryUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *primary) DescribeUser(r *restful.Request, w *restful.Response) {
	tk := r.Attribute(token.TOKEN_ATTRIBUTE_NAME).(*token.Token)

	req := user.NewDescriptUserRequestById(r.PathParameter("id"))

	ins, err := h.service.DescribeUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if tk.UserType != user.TYPE_SUPPER {
		if ins.Spec.Domain != tk.Domain || ins.Spec.Namespace != tk.Namespace {
			response.Failed(w, fmt.Errorf("no permission"))
			return
		}
	}

	response.Success(w, ins)
}

func init() {
	app.RegistryRESTfulApp(&primary{})
}
