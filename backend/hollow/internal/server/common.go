package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/selector"
)

func SkipRoutersMatcher() selector.MatchFunc {

	skipRouters := map[string]struct{}{
		"/user.v1.Users/PingConnect":   {}, //连接测试
		"/user.v1.Users/Login":         {}, //用户登录
		"/user.v1.Users/Register":      {}, //用户注册
		"/user.v1.Users/SendShortMsg":  {}, //发送短信
		"/forest.v1.Forests/GetForest": {}, //查询森林
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}
