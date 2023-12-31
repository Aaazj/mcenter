package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/Aaazj/mcenter/apps/namespace"
	//"github.com/Aaazj/mcenter/apps/policy"
	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/apps/token/provider"
	"github.com/Aaazj/mcenter/apps/token/security"
	"github.com/Aaazj/mcenter/conf"

	_ "github.com/Aaazj/mcenter/apps/token/provider/all"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	token.UnimplementedRPCServer
	log logger.Logger

	//policy  policy.Service
	ns      namespace.Service
	checker security.Checker
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	dc := db.Collection(s.Name())
	indexs := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "refresh_token", Value: bsonx.Int32(-1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "issue_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err = dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = dc

	s.log = zap.L().Named(s.Name())
	s.ns = app.GetInternalApp(namespace.AppName).(namespace.Service)
	//s.policy = app.GetInternalApp(policy.AppName).(policy.Service)

	s.checker, err = security.NewChecker()
	if err != nil {
		return fmt.Errorf("new checker error, %s", err)
	}

	// 初始化所有的auth provider
	if err := provider.Init(); err != nil {
		return err
	}

	return nil
}

func (s *service) Name() string {
	return token.AppName
}

func (s *service) Registry(server *grpc.Server) {
	token.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
