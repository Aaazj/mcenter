package impl

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/Aaazj/mcenter/apps/namespace"
	//"github.com/Aaazj/mcenter/apps/policy"
	//"github.com/Aaazj/mcenter/apps/role"
	"github.com/Aaazj/mcenter/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	namespace.UnimplementedRPCServer

	//role   role.Service
	//policy policy.Service
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	//i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())

	dc := db.Collection(i.Name())
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{
				{Key: "name", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err = dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	i.col = dc
	//i.role = app.GetInternalApp(role.AppName).(role.Service)
	//i.policy = app.GetInternalApp(policy.AppName).(policy.Service)
	return nil
}

func (i *impl) Name() string {
	return namespace.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	namespace.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
