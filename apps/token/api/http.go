package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/apps/user"
)

var (
	h = &handler{}
)

type handler struct {
	service token.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(token.AppName)
	h.service = app.GetInternalApp(token.AppName).(token.Service)
	return nil
}

func (h *handler) Name() string {
	return token.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"登录"}

	ws.Route(ws.POST("/").To(h.IssueToken))
	// .
	// 	Doc("颁发令牌").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(token.IssueTokenRequest{}).
	// 	Writes(token.Token{}).
	// 	Returns(200, "OK", token.Token{}))

	ws.Route(ws.DELETE("/").To(h.RevolkToken).
		Doc("撤销令牌").
		//Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/validate").To(h.ValidateToken).
		Doc("验证令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.ValidateTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	// 只有主账号才能查询
	ws.Route(ws.GET("/").To(h.QueryToken).
		Doc("令牌颁发记录").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, true).
		Metadata(label.Allow, user.TYPE_SUPPER).
		Reads(token.QueryTokenRequest{}).
		Writes(token.TokenSet{}).
		Returns(200, "OK", token.TokenSet{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
