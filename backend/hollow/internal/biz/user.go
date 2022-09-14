package biz

import (
	"context"

	v1 "hollow/api/hollow/v1"
	errors "hollow/internal/errors"
	models "hollow/internal/models"
	mfa "hollow/internal/pkg/middleware/MFA"
	utils "hollow/internal/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByID(ctx context.Context, userid int64) (*models.User, error)
	CreateUser(ctx context.Context, req *v1.RegisterUserRequest) error
	CheckIsUserExistByUsername(ctx context.Context, username string) bool
	CheckIsUserExistByID(ctx context.Context, userid int64) bool
	MFAGetQrCode(ctx context.Context) (string, string, error)
	MFAActivate(ctx context.Context, g *v1.MFAActivateRequest) error
	MFACancel(ctx context.Context, code string) error
	SendShortMsg(ctx context.Context, phone, code string) (data *models.ShortMsg, err error)
	CheckMsgCode(ctx context.Context, phone, code string) (result bool, err error)
	ReBindPhone(ctx context.Context, phone, code, mfacode string) error
	UpdateUserStatus(ctx context.Context, g *v1.UpdateUserStatusRequest) error
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) LoginUser(ctx context.Context, u *v1.LoginUserRequest) (*models.User, error) {

	data, err := uc.ur.GetUserByUsername(ctx, u.Username)

	if err != nil {
		return nil, err
	}

	// 被封禁
	if data.Status == 2 {
		return nil, errors.ErrUserBlocked
	}

	// MFA验证
	if data.Mfa_enabled {
		if len(u.Code) != 6 {
			return nil, errors.ErrNeedMFA
		}

		auth := mfa.NewGoogleAuth()
		ans, err := auth.VerifyCode(data.Mfa_secret, u.Code)

		if err != nil {
			return nil, err
		}

		if !ans {
			return nil, errors.ErrMFAVerifyFailed
		}
	}

	// 密码验证
	if utils.GenerateTokenSHA256(u.Password) != data.Password {
		return nil, errors.ErrUserCheckFailed
	}

	return data, err
}

func (uc *UserUsecase) RegisterUser(ctx context.Context, u *v1.RegisterUserRequest) (*models.User, error) {

	result, err := uc.ur.CheckMsgCode(ctx, u.Phone, u.Code)

	if err != nil {
		return nil, err
	}

	if !result {
		return nil, errors.ErrMsgCodeVerifyFailed
	}
	// 检查用户名是否重复
	if uc.ur.CheckIsUserExistByUsername(ctx, u.Username) {
		return nil, errors.ErrUserExisted
	}

	// 创建用户
	err = uc.ur.CreateUser(ctx, u)
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

func (uc *UserUsecase) SendShortMsg(ctx context.Context, u *v1.SendShortMsgRequest) (*models.ShortMsg, error) {
	code := utils.GetCapature() // 获取验证码
	data, err := uc.ur.SendShortMsg(ctx, u.Phone, code)

	if err != nil {
		return nil, errors.GenerateErrorNormal(err)
	}

	// 返回错误信息
	if data.Code != "OK" {
		return nil, errors.GenerateErrorString(data.Message)
	}

	return data, nil
}

func (uc *UserUsecase) ReBindPhone(ctx context.Context, u *v1.ReBindPhoneRequest) error {
	return uc.ur.ReBindPhone(ctx, u.Phone, u.Code, u.Mfacode)
}

func (uc *UserUsecase) GetUserInfo(ctx context.Context, u *v1.GetUserRequest) (*models.User, error) {
	return uc.ur.GetUserByID(ctx, u.Id)
}

func (uc *UserUsecase) MFAGetQRCode(ctx context.Context) (string, string, error) {
	return uc.ur.MFAGetQrCode(ctx)
}

func (uc *UserUsecase) MFAActivate(ctx context.Context, u *v1.MFAActivateRequest) error {
	return uc.ur.MFAActivate(ctx, u)
}

func (uc *UserUsecase) MFACancel(ctx context.Context, u *v1.MFACancelRequest) error {
	return uc.ur.MFACancel(ctx, u.Code)
}

func (uc *UserUsecase) UpdateUserStatus(ctx context.Context, u *v1.UpdateUserStatusRequest) error {
	return uc.ur.UpdateUserStatus(ctx, u)
}
