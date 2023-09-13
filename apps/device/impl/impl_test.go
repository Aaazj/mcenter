package impl_test

import (
	"context"

	"github.com/Aaazj/mcenter/apps/device"
	"github.com/Aaazj/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl device.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(device.AppName).(device.Service)
}
