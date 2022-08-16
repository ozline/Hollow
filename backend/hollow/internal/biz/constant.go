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
	ID         int64  //用户ID
	Status     int64  //用户状态
	Username   string //用户名
	Email      string //邮箱
	Phone      int64  //手机号
	Created_at int64  //创建时间
	Deleted_at int64  //删除时间
	Updated_at int64  //更新时间
	Password   string //密码
	Nickname   string //昵称
}

type Leaf struct {
	ID        int64  //消息ID
	Owner     int64  //所述对象
	Create_at int64  //创建时间
	Status    int64  //消息状态 0=匿名 1=实名 2=封禁
	Message   string //消息
}
