package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/Aaazj/mcenter/apps/code/api"
	_ "github.com/Aaazj/mcenter/apps/domain/api"
	_ "github.com/Aaazj/mcenter/apps/endpoint/api"

	//_ "github.com/Aaazj/mcenter/apps/health/api"
	_ "github.com/Aaazj/mcenter/apps/instance/api"
	_ "github.com/Aaazj/mcenter/apps/resource/api"
	_ "github.com/Aaazj/mcenter/apps/service/api"

	//_ "github.com/Aaazj/mcenter/apps/setting/api"
	_ "github.com/Aaazj/mcenter/apps/token/api"
	_ "github.com/Aaazj/mcenter/apps/user/api"
)
