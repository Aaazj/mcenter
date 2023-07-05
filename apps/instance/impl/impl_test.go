package impl_test

import (
	"context"

	"github.com/Aaazj/mcenter/apps/instance"
	"github.com/Aaazj/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl instance.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(instance.AppName).(instance.Service)
}
