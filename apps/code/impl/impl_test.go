package impl_test

import (
	"context"

	"github.com/Aaazj/mcenter/apps/code"
	"github.com/Aaazj/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl code.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(code.AppName).(code.Service)
}
