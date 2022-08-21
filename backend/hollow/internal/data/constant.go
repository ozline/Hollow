package data

import (
	"context"
	"hollow/internal/pkg/middleware/auth"
)

//TABLES
const (
	TABLE_USERS   = "users"
	TABLE_FOREST  = "forest"
	TABLE_COMMENT = "comment"
)

func GetUserInfo(ctx context.Context) *auth.CurrentUser {
	currentUser := auth.FromContext(ctx)
	return currentUser
}
