package biz

import (
	"context"

	v1 "hollow/api/hollow/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound  = errors.New(422, v1.ErrorReason_INFOMATION_NOT_FOUND.String(), "user not found")
	ErrMissingParams = errors.New(422, v1.ErrorReason_PARAMS_ILLEGAL.String(), "missing params data")
	ErrParamsIllegal = errors.New(422, v1.ErrorReason_PARAMS_ILLEGAL.String(), "username of password invalid")
	ErrUserExisted   = errors.New(422, v1.ErrorReason_PARAMS_ILLEGAL.String(), "user existed")
)

type UserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, req *v1.RegisterUserRequest) error
	CheckIsUserExist(ctx context.Context, username string) bool
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

type User struct {
	ID        int64
	Status    int64
	Username  string
	Email     string
	Phone     int64
	Create_at int64
	Delete_at int64
	Update_at int64
	Password  string
	Nickname  string
	Token     string
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) LoginUser(ctx context.Context, u *v1.LoginUserRequest) (*User, error) {
	if len(u.Username) == 0 || len(u.Password) == 0 {
		return nil, ErrMissingParams
	}
	data, err := uc.ur.GetUserByUsername(ctx, u.Username)
	if err != nil {
		return nil, err
	}
	if GenerateTokenSHA256(u.Password) != data.Password {
		return nil, ErrParamsIllegal
	}

	return data, err
}

func (uc *UserUsecase) RegisterUser(ctx context.Context, u *v1.RegisterUserRequest) (*User, error) {
	if len(u.Username) == 0 || len(u.Password) == 0 || u.Phone == 0 || len(u.Email) == 0 {
		return nil, errors.New(422, v1.ErrorReason_PARAMS_ILLEGAL.String(), "missing params data")
	}

	//检查用户名是否重复
	if uc.ur.CheckIsUserExist(ctx, u.Username) {
		return nil, ErrUserExisted
	}

	//创建用户
	err := uc.ur.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	//注册成功后调用一次获取用户信息的API，回调用户信息
	data, err := uc.ur.GetUserByUsername(ctx, u.Username)

	if err != nil {
		return nil, err
	}

	return data, nil
}
