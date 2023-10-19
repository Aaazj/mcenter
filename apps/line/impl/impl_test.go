package impl_test

import (
	"context"

	"github.com/Aaazj/mcenter/apps/line"
	"github.com/Aaazj/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl line.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(line.AppName).(line.Service)
}
