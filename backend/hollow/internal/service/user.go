package service

import (
	"context"
	"strconv"

	v1 "hollow/api/hollow/v1"
	"hollow/internal/biz"
	"hollow/internal/pkg/middleware/auth"

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

//用户登录
func (s *UserService) Login(ctx context.Context, req *v1.LoginUserRequest) (reply *v1.LoginUserReply, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		return nil, ErrMissingParams
	}

	data, err := s.uc.LoginUser(ctx, req)

	if err != nil {
		return nil, err
	}

	token, _ := auth.GetAuthToken(strconv.FormatInt(data.ID, 10), data.Username, int(data.Status), "MTAxNTkwMTg1Mw==")
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
			CreatedAt: data.Create_at,
		},
	}, nil
}

//用户注册
func (s *UserService) Register(ctx context.Context, req *v1.RegisterUserRequest) (reply *v1.RegisterUserReply, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 || req.Phone == 0 || len(req.Email) == 0 {
		return nil, ErrMissingParams
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
			CreatedAt: data.Create_at,
		},
	}, nil
}
