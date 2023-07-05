package impl_test

import (
	"context"

	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl token.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(token.AppName).(token.Service)
}
