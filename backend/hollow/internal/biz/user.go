package biz

import (
	"context"

	v1 "hollow/api/hollow/v1"
	"hollow/internal/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, req *v1.RegisterUserRequest) error
	CheckIsUserExistByUsername(ctx context.Context, username string) bool
	CheckIsUserExistByID(ctx context.Context, userid int64) bool
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) LoginUser(ctx context.Context, u *v1.LoginUserRequest) (*User, error) {

	data, err := uc.ur.GetUserByUsername(ctx, u.Username)
	if err != nil {
		return nil, err
	}
	if utils.GenerateTokenSHA256(u.Password) != data.Password {
		return nil, ErrUserCheckFailed
	}

	return data, err
}

func (uc *UserUsecase) RegisterUser(ctx context.Context, u *v1.RegisterUserRequest) (*User, error) {
	//检查用户名是否重复
	if uc.ur.CheckIsUserExistByUsername(ctx, u.Username) {
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
