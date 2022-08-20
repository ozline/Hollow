package data

import (
	"context"
	v1 "hollow/api/hollow/v1"

	"hollow/internal/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/errors"
)

//TABLES
const (
	TABLE_USERS  = "users"
	TABLE_FOREST = "forest"
)

//ERRORS
var (
	//UNIVERSE
	ErrNormal = errors.New(500, v1.ErrorReason_NORMAL_ERROR.String(), "Unknown Error")

	//USERS
	ErrUserNotExisted = errors.New(422, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "User Not Existed")

	//FOREST
	ErrUserInvalid = errors.New(422, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "User Invalid")
)

func getUserInfo(ctx context.Context) *auth.CurrentUser {
	currentUser := auth.FromContext(ctx)
	return currentUser
}
