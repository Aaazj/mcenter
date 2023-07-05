package impl_test

import (
	"context"

	"github.com/Aaazj/mcenter/apps/policy"
	"github.com/Aaazj/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(policy.AppName).(policy.Service)
}
