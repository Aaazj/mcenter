package impl_test

import (
	"context"

	"github.com/Aaazj/mcenter/apps/domain"
	"github.com/Aaazj/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl domain.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(domain.AppName).(domain.Service)
}
