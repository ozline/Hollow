package data

import (
	"context"

	v1 "hollow/api/hollow/v1"
	"hollow/internal/biz"
	"hollow/internal/pkg/utils"

	errors "hollow/internal/errors"
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

func (r *userRepo) CreateUser(ctx context.Context, g *v1.RegisterUserRequest) error {

	timeStamp := utils.GetTimestamp13()
	u := types.User{
		Username:   g.Username,
		Password:   utils.GenerateTokenSHA256(g.Password),
		Phone:      g.Phone,
		Email:      g.Email,
		Created_at: timeStamp,
		Updated_at: 0,
		Nickname:   g.Username,
	}

	res := r.data.db.Table(TABLE_USERS).Create(&u)

	return res.Error
}
