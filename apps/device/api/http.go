package api

import (
	"github.com/Aaazj/mcenter/apps/device"

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
	service device.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(device.AppName)
	h.service = app.GetInternalApp(device.AppName).(device.Service)
	return nil
}

func (h *handler) Name() string {
	return device.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"设备资源管理"}
	ws.Route(ws.POST("/").To(h.AllocationDevice).
		Doc("分配资源").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(device.AllocationRequest{}))

	ws.Route(ws.GET("/").To(h.QueryDevice).
		Doc("查询服务功能列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list"))

	ws.Route(ws.GET("/{name}").To(h.DescribeDevice).
		Doc("查询服务功能详情").
		Param(ws.PathParameter("name", "name of the device").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/").To(h.ReleaseDevices).
		Doc("释放资源").
		Metadata(restfulspec.KeyOpenAPITags, tags))

}

func init() {
	app.RegistryRESTfulApp(h)
}
