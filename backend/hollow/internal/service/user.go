package service

import (
	"context"

	v1 "hollow/api/hollow/v1"
	"hollow/internal/biz"
	"hollow/internal/pkg/middleware/auth"

	errors "hollow/internal/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUsersServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// 用户登录
func (s *UserService) Login(ctx context.Context, req *v1.LoginUserRequest) (reply *v1.LoginUserReply, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		return nil, errors.ErrMissingParams
	}

	data, err := s.uc.LoginUser(ctx, req)

	if err != nil {
		return nil, err
	}

	token, _ := auth.GetAuthToken(data.ID, data.Username, data.Status, "MTAxNTkwMTg1Mw==")
	return &v1.LoginUserReply{
		Code:  200,
		Msg:   "ok",
		Token: token,
		Data: &v1.User{
			Userid:    data.ID,
			Phone:     data.Phone,
			Email:     data.Email,
			Username:  data.Username,
			Nickname:  data.Nickname,
			Status:    data.Status,
			CreatedAt: data.Created_at,
		},
	}, nil
}

// 用户注册
func (s *UserService) Register(ctx context.Context, req *v1.RegisterUserRequest) (reply *v1.RegisterUserReply, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 || req.Phone == 0 || len(req.Email) == 0 {
		return nil, errors.ErrMissingParams
	}

	data, err := s.uc.RegisterUser(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.RegisterUserReply{
		Code: 200,
		Msg:  "ok",
		Data: &v1.User{
			Userid:    data.ID,
			Phone:     data.Phone,
			Email:     data.Email,
			Username:  data.Username,
			Nickname:  data.Nickname,
			Status:    data.Status,
			CreatedAt: data.Created_at,
		},
	}, nil
}

// 用户注册
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (reply *v1.GetUserReply, err error) {

	data, err := s.uc.GetUserInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.GetUserReply{
		Code: 200,
		Msg:  "ok",
		Data: &v1.User{
			Userid:    data.ID,
			Phone:     data.Phone,
			Email:     data.Email,
			Username:  data.Username,
			Nickname:  data.Nickname,
			Status:    data.Status,
			CreatedAt: data.Created_at,
		},
	}, nil
}
