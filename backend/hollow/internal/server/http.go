package server

import (
	"context"
	"fmt"
	v1 "hollow/api/hollow/v1"
	"hollow/internal/conf"
	"hollow/internal/pkg/middleware/auth"
	"hollow/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func SkipRoutersMatcherHTTP() selector.MatchFunc {

	skipRouters := map[string]struct{}{
		"/user.v1.Users/Login":    {}, //用户登录
		"/user.v1.Users/Register": {}, //用户注册
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			// return false
			return false
		}
		fmt.Println(operation)
		return false
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, jwtc *conf.Auth, users *service.UserService, forests *service.ForestService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			selector.Server(auth.JWTAuth(jwtc.Secret)).Match(SkipRoutersMatcherHTTP()).Build(),
			//异常恢复
			recovery.Recovery(
				recovery.WithLogger(log.DefaultLogger),
				recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {
					return nil
				}),
			),

			//参数校验
			validate.Validator(),

			//限流器
			ratelimit.Server(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterUsersHTTPServer(srv, users)
	v1.RegisterForestsHTTPServer(srv, forests)
	return srv
}
