package api

import (
	"github.com/Aaazj/mcenter/apps/instance"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service instance.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(instance.AppName)
	h.service = app.GetInternalApp(instance.AppName).(instance.Service)
	return nil
}

func (h *handler) Name() string {
	return instance.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"服务实例管理"}

	ws.Route(ws.GET("/").To(h.SearchInstance).
		Doc("搜索实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(instance.RegistryRequest{}).
		Writes(instance.Instance{}))

	ws.Route(ws.POST("/").To(h.RegistryInstance).
		Doc("实例注册").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(instance.UnregistryRequest{}).
		Writes(instance.Instance{}))

	ws.Route(ws.DELETE("/{instance_id}").To(h.RegistryInstance).
		Doc("实例注销").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(instance.SearchRequest{}).
		Writes(instance.InstanceSet{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
