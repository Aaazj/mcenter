package rpc

import (
	"context"
	"sync"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Aaazj/mcenter/apps/audit"
	"github.com/Aaazj/mcenter/apps/device"
	"github.com/Aaazj/mcenter/apps/endpoint"
	"github.com/Aaazj/mcenter/apps/image"
	"github.com/Aaazj/mcenter/apps/instance"
	"github.com/Aaazj/mcenter/apps/line"
	"github.com/Aaazj/mcenter/apps/namespace"

	"github.com/Aaazj/mcenter/apps/service"
	"github.com/Aaazj/mcenter/apps/token"
	"github.com/Aaazj/mcenter/apps/user"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

// NewClient todo
func NewClient(conf *Config) (*ClientSet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), conf.Timeout())
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		// mcenter服务地址
		conf.Address,
		// 不使用TLS
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 开启认证
		grpc.WithPerRPCCredentials(conf.Credentials()),
		// gprc 支持的负载均衡策略: https://github.com/grpc/grpc/blob/master/doc/load-balancing.md
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		// 将异常转化为 API Exception
		grpc.WithChainUnaryInterceptor(NewExceptionUnaryClientInterceptor().UnaryClientInterceptor),
		// Grpc Trace
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conf: conf,
		conn: conn,
		log:  zap.L().Named("mcenter.rpc"),
	}, nil
}

// Client 客户端
type ClientSet struct {
	conf *Config
	conn *grpc.ClientConn
	log  logger.Logger
	svr  *service.Service
	lock sync.Mutex
}

// 返回客户端服务信息
func (c *ClientSet) ClientInfo(ctx context.Context) (*service.Service, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.svr != nil {
		return c.svr, nil
	}

	req := service.NewValidateCredentialRequest(c.conf.ClientID, c.conf.ClientSecret)
	svc, err := c.Service().ValidateCredential(ctx, req)
	if err != nil {
		return nil, err
	}
	c.svr = svc
	return c.svr, nil
}

// Instance服务的SDK
func (c *ClientSet) Health() healthgrpc.HealthClient {
	return healthgrpc.NewHealthClient(c.conn)
}

// Instance服务的SDK
func (c *ClientSet) Instance() instance.RPCClient {
	return instance.NewRPCClient(c.conn)
}

// Service服务的SDK
func (c *ClientSet) Service() service.RPCClient {
	return service.NewRPCClient(c.conn)
}

// Token服务的SDK
func (c *ClientSet) Token() token.RPCClient {
	return token.NewRPCClient(c.conn)
}

// Namespace服务的SDK
func (c *ClientSet) Namespace() namespace.RPCClient {
	return namespace.NewRPCClient(c.conn)
}

// User服务的SDK
func (c *ClientSet) User() user.RPCClient {
	return user.NewRPCClient(c.conn)
}

// Device服务的SDK
func (c *ClientSet) Device() device.RPCClient {
	return device.NewRPCClient(c.conn)
}

// Endpoint服务的SDK
func (c *ClientSet) Endpoint() endpoint.RPCClient {
	return endpoint.NewRPCClient(c.conn)
}

// Audit服务的SDK
func (c *ClientSet) Audit() audit.RPCClient {
	return audit.NewRPCClient(c.conn)
}

func (c *ClientSet) Image() image.RPCClient {
	return image.NewRPCClient(c.conn)
}

func (c *ClientSet) Line() line.RPCClient {
	return line.NewRPCClient(c.conn)
}
