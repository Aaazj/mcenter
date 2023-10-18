package api

import (
	"github.com/Aaazj/mcenter/apps/namespace"
	"github.com/Aaazj/mcenter/apps/user"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service namespace.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(namespace.AppName)
	h.service = app.GetInternalApp(namespace.AppName).(namespace.Service)
	return nil
}

func (h *handler) Name() string {
	return namespace.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"域管理"}

	ws.Route(ws.POST("/").To(h.CreateNamespace).
		Doc("创建空间").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Allow, user.TYPE_SUPPER).
		Metadata(label.Audit, true).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, "Create").
		Reads(namespace.CreateNamespaceRequest{}).
		Writes(namespace.Namespace{}))

	ws.Route(ws.GET("/{d_id}/{id}").To(h.DescribeNamespace).
		Doc("查询空间").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Allow, user.TYPE_SUPPER).
		Writes(namespace.Namespace{}).
		Returns(200, "OK", namespace.Namespace{}))

	ws.Route(ws.DELETE("/{d_id}/{id}").To(h.DeleteNamespace).
		Doc("删除空间").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Audit, true).
		Metadata(label.Action, "Delete").
		Metadata(label.Allow, user.TYPE_SUPPER))

	ws.Route(ws.GET("/").To(h.QueryNamespace).
		Doc("查询空间列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Allow, user.TYPE_SUPPER).
		Writes(namespace.NamespaceSet{}).
		Returns(200, "OK", namespace.NamespaceSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{d_id}/{id}").To(h.PutNamespace).
		Metadata(label.Auth, true).
		Metadata(label.Allow, user.TYPE_SUPPER).
		Doc("修改子账号").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(namespace.CreateNamespaceRequest{}))

}

func init() {
	app.RegistryRESTfulApp(h)
}
