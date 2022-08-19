package server

import (
	"context"
	v1 "hollow/api/hollow/v1"
	"hollow/internal/conf"
	"hollow/internal/pkg/middleware/auth"
	"hollow/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewSkipRoutersMatcher() selector.MatchFunc {

	skipRouters := map[string]struct{}{
		"/user.v1.Users/Login":    {}, //用户登录
		"/user.v1.Users/Register": {}, //用户注册
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, jwtc *conf.Auth, users *service.UserService, forests *service.ForestService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			mmd.Server(),
			selector.Server(auth.JWTAuth(jwtc.Secret)).Match(NewSkipRoutersMatcher()).Build(),
			recovery.Recovery(
				recovery.WithLogger(log.DefaultLogger),
				recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {

					//还没遇到需要处理的panic..
					return nil
				}),
			),
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUsersServer(srv, users)
	v1.RegisterForestsServer(srv, forests)
	return srv
}
