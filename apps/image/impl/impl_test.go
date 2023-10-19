package impl_test

import (
	"context"

	"github.com/Aaazj/mcenter/apps/image"
	"github.com/Aaazj/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl image.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(image.AppName).(image.Service)
}
