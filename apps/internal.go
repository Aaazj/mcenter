package all

import (
	// 注册所有内部服务模块, 无须对外暴露的服务, 用于内部依赖
	//_ "github.com/Aaazj/mcenter/apps/counter/impl"
	//_ "github.com/Aaazj/mcenter/apps/ip2region/impl"
	//_ "github.com/Aaazj/mcenter/apps/setting/impl"
	//_ "github.com/Aaazj/mcenter/apps/storage/impl"

	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	//_ "github.com/Aaazj/mcenter/apps/code/impl"
	_ "github.com/Aaazj/mcenter/apps/domain/impl"
	_ "github.com/Aaazj/mcenter/apps/endpoint/impl"

	//_ "github.com/Aaazj/mcenter/apps/health/impl"
	_ "github.com/Aaazj/mcenter/apps/instance/impl"
	_ "github.com/Aaazj/mcenter/apps/namespace/impl"

	//_ "github.com/Aaazj/mcenter/apps/notify/impl"
	//_ "github.com/Aaazj/mcenter/apps/policy/impl"
	_ "github.com/Aaazj/mcenter/apps/resource/impl"
	//_ "github.com/Aaazj/mcenter/apps/role/impl"
	_ "github.com/Aaazj/mcenter/apps/audit/impl"
	_ "github.com/Aaazj/mcenter/apps/device/impl"
	_ "github.com/Aaazj/mcenter/apps/service/impl"
	_ "github.com/Aaazj/mcenter/apps/token/impl"
	_ "github.com/Aaazj/mcenter/apps/user/impl"
)
