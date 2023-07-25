package tools

import (
	"os"

	"github.com/Aaazj/mcenter/conf"
	//"github.com/Aaazj/mcenter/tracer"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/cache/memory"
	"github.com/infraboard/mcube/logger/zap"

	// 注册所有服务
	_ "github.com/Aaazj/mcenter/apps"
)

func DevelopmentSetup() {
	// 初始化日志实例
	zap.DevelopmentSetup()

	// 初始化配置, 提前配置好/etc/unit_test.env
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	//tracer.InitTracer()

	// 初始化缓存
	ins := memory.NewCache(conf.C().Cache.Memory)
	cache.SetGlobal(ins)

	// 初始化全局app
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
}

func AccessToken() string {
	return os.Getenv("MCENTER_ACCESS_TOKEN")
}
