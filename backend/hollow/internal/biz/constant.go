package biz

import (
	v1 "hollow/api/hollow/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

//ERRORS
var (

	//USERS
	ErrUserCheckFailed = errors.New(422, v1.ErrorReason_INFORMATION_ILLEGAL.String(), "user check failed")
	ErrUserNotFound    = errors.New(422, v1.ErrorReason_INFOMATION_NOT_FOUND.String(), "user not found")
	ErrUserExisted     = errors.New(422, v1.ErrorReason_PARAMS_ILLEGAL.String(), "user existed")
)

//TYPES
type User struct {
	ID         int64
	Status     int64
	Username   string
	Email      string
	Phone      int64
	Created_at int64
	Deleted_at int64
	Updated_at int64
	Password   string
	Nickname   string
}

type Leaf struct {
	ID        int64
	Owner     int64
	Create_at int64
	Status    int64
	Message   string
}
