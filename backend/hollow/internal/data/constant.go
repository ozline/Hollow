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

//TYPES
type User struct {
	ID         int64
	Username   string
	Password   string
	Status     int64
	Email      string
	Nickname   string
	Phone      int64
	Updated_at int64
	Created_at int64
}
type Leaf struct {
	ID        int64
	Owner     int64
	Create_at int64
	Status    int64
	Message   string
}
