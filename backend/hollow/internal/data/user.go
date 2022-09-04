package data

import (
	"context"
	"strconv"

	v1 "hollow/api/hollow/v1"
	"hollow/internal/biz"
	"hollow/internal/pkg/utils"

	errors "hollow/internal/errors"
	mfa "hollow/internal/pkg/middleware/MFA"
	types "hollow/internal/types"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CheckIsUserExistByUsername(ctx context.Context, username string) bool {

	var count int64
	_ = r.data.db.Table(TABLE_USERS).Where("username = ?", username).Limit(1).Count(&count)
	return count != 0
}

func (r *userRepo) CheckIsUserExistByID(ctx context.Context, userid int64) bool {
	var count int64
	_ = r.data.db.Table(TABLE_USERS).Where("id = ?", userid).Limit(1).Count(&count)
	return count != 0
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (user *types.User, err error) {
	u := new(types.User)
	var count int64
	res := r.data.db.Table(TABLE_USERS).Where("username = ?", username).Count(&count)
	if res.Error != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.ErrUserNotExisted
	}

	_ = res.First(&u)

	return u, nil
}

func (r *userRepo) GetUserByID(ctx context.Context, userid int64) (user *types.User, err error) {
	u := new(types.User)
	var count int64
	res := r.data.db.Table(TABLE_USERS).Where("id = ?", userid).Count(&count)
	if res.Error != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.ErrUserNotExisted
	}

	_ = res.First(&u)

	return u, nil
}

func (r *userRepo) CreateUser(ctx context.Context, g *v1.RegisterUserRequest) error {

	num, err := strconv.ParseInt(g.Phone, 10, 64)

	if err != nil {
		return err
	}

	timeStamp := utils.GetTimestamp13()
	u := types.User{
		ID:         GetSnowflakeID(r.data.node),
		Username:   g.Username,
		Password:   utils.GenerateTokenSHA256(g.Password),
		Phone:      num,
		Email:      g.Email,
		Created_at: timeStamp,
		Updated_at: 0,
		Nickname:   g.Username,
	}

	res := r.data.db.Table(TABLE_USERS).Create(&u)

	return res.Error
}

func (r *userRepo) MFAGetQrCode(ctx context.Context) (string, string, error) {
	user := GetUserInfo(ctx)

	u := new(types.User)

	res := r.data.db.Table(TABLE_USERS).Where("id = ?", user.ID).First(&u)

	if res.Error != nil {
		return "", "", res.Error
	}

	if u.Mfa_enabled {
		return "", "", errors.ErrHaveActivatedMFA
	}

	// 获取秘钥
	auth := mfa.NewGoogleAuth()
	secret := auth.GetSecret()

	return auth.GetQrcodeUrl(u.Email, secret), secret, nil
}

// 激活MFA
func (r *userRepo) MFAActivate(ctx context.Context, g *v1.MFAActivateRequest) error {
	user := GetUserInfo(ctx)

	u := new(types.User)

	res := r.data.db.Table(TABLE_USERS).Where("id = ?", user.ID).First(&u)

	if res.Error != nil {
		return res.Error
	}

	if u.Mfa_enabled {
		return errors.ErrHaveActivatedMFA
	}

	// 启用MFA
	auth := mfa.NewGoogleAuth()
	ans, err := auth.VerifyCode(g.Secret, g.Code)

	if err != nil {
		return err
	}

	if !ans {
		return errors.ErrMFAVerifyFailed
	}

	u.Mfa_enabled = true
	u.Mfa_secret = g.Secret

	res.Save(&u)

	return nil
}

// 解绑MFA
func (r *userRepo) MFACancel(ctx context.Context, code string) error {
	user := GetUserInfo(ctx)

	u := new(types.User)

	res := r.data.db.Table(TABLE_USERS).Where("id = ?", user.ID).First(&u)

	if res.Error != nil {
		return res.Error
	}

	if !u.Mfa_enabled {
		return errors.ErrNotEnabledMFA
	}

	auth := mfa.NewGoogleAuth()
	ans, err := auth.VerifyCode(u.Mfa_secret, code)

	if err != nil {
		return err
	}

	if !ans {
		return errors.ErrMFAVerifyFailed
	}

	u.Mfa_enabled = false
	u.Mfa_secret = ""
	res.Save(&u)

	return nil
}

// 发送短信
func (r *userRepo) SendShortMsg(ctx context.Context, phone, code string) (data *types.ShortMsg, err error) {
	rs, err := r.data.dbRedis.Do("TTL", phone)

	if err != nil {
		return nil, err
	}

	// 存在的话说明还没过期
	if utils.Intval(rs) >= 0 {
		return nil, errors.GenerateErrorString("发信时间间隔过短,还需" + utils.Strval(rs) + "秒")
	}

	_, err = r.data.dbRedis.Do("SET", phone, code)

	if err != nil {
		return nil, err
	}

	// 设置5分钟过期
	rs, err = r.data.dbRedis.Do("EXPIRE", phone, 60) // 设置1分钟过期

	if err != nil {
		return nil, err
	}

	if utils.Strval(rs) != "1" {
		return nil, errors.ErrNormal
	}

	return r.data.shortmsg.PostMsg(phone, code) //开始发信
}

func (r *userRepo) CheckMsgCode(ctx context.Context, phone, code string) (result bool, err error) {
	rs, err := r.data.dbRedis.Do("GET", phone)

	if utils.Strval(rs) == "" {
		return false, errors.ErrMsgCodeInvalid
	}

	if err != nil {
		return false, err
	}

	if utils.Strval(rs) != code {
		return false, errors.ErrMsgCodeVerifyFailed
	}

	return true, nil
}

func (r *userRepo) ReBindPhone(ctx context.Context, phone, code, mfacode string) error {
	user := GetUserInfo(ctx)

	u := new(types.User)

	res := r.data.db.Table(TABLE_USERS).Where("id = ?", user.ID).First(&u)

	if res.Error != nil {
		return res.Error
	}

	// 如果开启MFA，则优先验证MFA
	if u.Mfa_enabled {
		auth := mfa.NewGoogleAuth()
		ans, err := auth.VerifyCode(u.Mfa_secret, mfacode)

		if err != nil {
			return err
		}

		if !ans {
			return errors.ErrMFAVerifyFailed
		}
	}

	result, err := r.CheckMsgCode(ctx, phone, code)

	if err != nil {
		return err
	}

	if !result {
		return errors.ErrMsgCodeVerifyFailed
	}

	u.Phone, err = strconv.ParseInt(phone, 10, 64)

	if err != nil {
		return err
	}

	res.Save(&u)

	return nil
}

func (r *userRepo) UpdateUserStatus(ctx context.Context, g *v1.UpdateUserStatusRequest) error {

	user := GetUserInfo(ctx)

	if user.Status != 1 {
		return errors.ErrUserInsufficientPermissions
	}

	u := new(types.User)

	res := r.data.db.Table(TABLE_USERS).Where("id = ?", g.Id).First(&u)

	if res.Error != nil {
		return res.Error
	}

	u.Status = g.Status

	res.Save(&u)

	return res.Error
}
