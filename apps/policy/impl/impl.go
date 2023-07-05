package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/Aaazj/mcenter/apps/endpoint"
	"github.com/Aaazj/mcenter/apps/namespace"
	"github.com/Aaazj/mcenter/apps/policy"
	"github.com/Aaazj/mcenter/apps/role"
	"github.com/Aaazj/mcenter/apps/user"
	"github.com/Aaazj/mcenter/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	policy.UnimplementedRPCServer

	user      user.Service
	role      role.Service
	namespace namespace.Service
	endpoint  endpoint.Service
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())

	i.user = app.GetInternalApp(user.AppName).(user.Service)
	i.role = app.GetInternalApp(role.AppName).(role.Service)
	i.namespace = app.GetInternalApp(namespace.AppName).(namespace.Service)
	i.endpoint = app.GetInternalApp(endpoint.AppName).(endpoint.Service)
	return nil
}

func (i *impl) Name() string {
	return policy.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	policy.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
