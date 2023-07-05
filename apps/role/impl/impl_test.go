package impl_test

import (
	"context"

	"github.com/Aaazj/mcenter/apps/role"
	"github.com/Aaazj/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl role.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(role.AppName).(role.Service)
}
