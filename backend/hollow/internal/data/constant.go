package data

import (
	v1 "hollow/api/hollow/v1"

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
	ErrNormal = errors.New(500, v1.ErrorReason_NORMAL_ERROR.String(), "unknow error")

	//USERS
	ErrUserNotExisted = errors.New(422, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "user not existed")

	//FOREST
)
