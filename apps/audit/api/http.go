package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Aaazj/mcenter/apps/audit"
)

var (
	h = &handler{}
)

type handler struct {
	service audit.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(audit.AppName)
	h.service = app.GetInternalApp(audit.AppName).(audit.Service)
	return nil
}

func (h *handler) Name() string {
	return audit.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	//tags := []string{"审计管理"}

	ws.Route(ws.GET("/{id}").To(h.DescribeAudit).
		Doc("查询域"))

	ws.Route(ws.GET("/").To(h.QueryAudit).
		Doc("查询空间列表"))
}

func init() {
	app.RegistryRESTfulApp(h)
}
