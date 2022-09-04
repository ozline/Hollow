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

func (s *UserService) PingConnect(ctx context.Context, req *v1.PingConnectRequest) (reply *v1.PingConnectReply, err error) {
	return &v1.PingConnectReply{
		Code: 200,
		Msg:  "ok",
	}, nil
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
			Userid:     data.ID,
			Phone:      data.Phone,
			Email:      data.Email,
			Username:   data.Username,
			Nickname:   data.Nickname,
			Status:     data.Status,
			CreatedAt:  data.Created_at,
			MfaEnabled: data.Mfa_enabled,
		},
	}, nil
}

// 发送手机验证码
func (s *UserService) SendShortMsg(ctx context.Context, req *v1.SendShortMsgRequest) (reply *v1.SendShortMsgReply, err error) {
	data, err := s.uc.SendShortMsg(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.SendShortMsgReply{
		Code: 200,
		Msg:  "ok",
		Data: &v1.ShortMsg{
			Code:      data.Code,
			Message:   data.Message,
			Bizid:     data.BizId,
			Requestid: data.RequestId,
		},
	}, nil
}

// 换绑手机号
func (s *UserService) ReBindPhone(ctx context.Context, req *v1.ReBindPhoneRequest) (reply *v1.ReBindPhoneReply, err error) {
	err = s.uc.ReBindPhone(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.ReBindPhoneReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

// 用户注册
func (s *UserService) Register(ctx context.Context, req *v1.RegisterUserRequest) (reply *v1.RegisterUserReply, err error) {

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

// 获取用户信息
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (reply *v1.GetUserReply, err error) {

	data, err := s.uc.GetUserInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.GetUserReply{
		Code: 200,
		Msg:  "ok",
		Data: &v1.User{
			Userid:     data.ID,
			Phone:      data.Phone,
			Email:      data.Email,
			Username:   data.Username,
			Nickname:   data.Nickname,
			Status:     data.Status,
			CreatedAt:  data.Created_at,
			MfaEnabled: data.Mfa_enabled,
		},
	}, nil
}

func (s *UserService) MFAGetQRCode(ctx context.Context, req *v1.NullRequest) (reply *v1.MFAGetQRCodeReply, err error) {

	qrlink, secret, err := s.uc.MFAGetQRCode(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.MFAGetQRCodeReply{
		Code: 200,
		Msg:  "ok",
		Data: &v1.MFA{
			Qrlink: qrlink,
			Secret: secret,
		},
	}, nil
}

// 激活MFA
func (s *UserService) MFAActivate(ctx context.Context, req *v1.MFAActivateRequest) (reply *v1.MFAActivateReply, err error) {

	err = s.uc.MFAActivate(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.MFAActivateReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

// 解绑MFA
func (s *UserService) MFACancel(ctx context.Context, req *v1.MFACancelRequest) (reply *v1.MFACancelReply, err error) {

	err = s.uc.MFACancel(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.MFACancelReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

// 更新用户状态
func (s *UserService) UpdateUserStatus(ctx context.Context, req *v1.UpdateUserStatusRequest) (reply *v1.UpdateUserStatusReply, err error) {

	err = s.uc.UpdateUserStatus(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.UpdateUserStatusReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}
