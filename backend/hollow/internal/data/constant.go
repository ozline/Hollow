package data

import (
	"context"
	"hollow/internal/pkg/middleware/auth"

	snowflake "github.com/bwmarrin/snowflake"
)

// TABLES
const (
	TABLE_USERS   = "users"
	TABLE_FOREST  = "forest"
	TABLE_COMMENT = "comment"
	TABLE_LIKE    = "likes"
	TABLE_REPORT  = "reports"
)

// 获取用户信息
func GetUserInfo(ctx context.Context) *auth.CurrentUser {
	currentUser := auth.FromContext(ctx)
	return currentUser
}

// 获取雪花ID
func GetSnowflakeID(node *snowflake.Node) int64 {
	return node.Generate().Int64()
}
