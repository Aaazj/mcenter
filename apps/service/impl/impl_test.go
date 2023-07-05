package impl_test

import (
	"context"

	"github.com/infraboard/mcube/app"

	// 注册所有服务
	"github.com/Aaazj/mcenter/apps/service"
	"github.com/Aaazj/mcenter/test/tools"
)

var (
	impl service.MetaService
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(service.AppName).(service.MetaService)
}
