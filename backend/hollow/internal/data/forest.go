package data

import (
	"context"
	"fmt"

	v1 "hollow/api/hollow/v1"
	"hollow/internal/biz"

	errors "hollow/internal/errors"
	utils "hollow/internal/pkg/utils"
	types "hollow/internal/types"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type forestRepo struct {
	data *Data
	log  *log.Helper
}

func NewForestRepo(data *Data, logger log.Logger) biz.ForestRepo {
	return &forestRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *forestRepo) PushLeaf(ctx context.Context, g *v1.PushLeafRequest) error {
	timeStamp := utils.GetTimestamp13()
	user := GetUserInfo(ctx) //获取用户信息

	if user.ID == -1 {
		return errors.ErrUserInvalid
	}

	u := types.Leaf{
		Owner:     user.ID,
		Create_at: timeStamp,
		Status:    g.Status,
		Message:   g.Message,
	}

	fmt.Println("叶子信息:", u)

	res := r.data.db.Table(TABLE_FOREST).Create(&u)

	return res.Error
}

func (r *forestRepo) CommentLeaf(ctx context.Context, g *v1.CommentLeafRequest) error {
	timeStamp := utils.GetTimestamp13()
	user := GetUserInfo(ctx)

	if user.ID == -1 {
		return errors.ErrUserInvalid
	}

	if g.Father != 0 { //验证用户是否存在，不存在则直接返回
		var count int64
		_ = r.data.db.Table(TABLE_USERS).Where("id = ?", g.Father).Limit(1).Count(&count)
		if count == 0 {
			return errors.ErrFatherNotExistd
		}
	}

	u := types.Comment{
		Owner:      user.ID,
		Root:       g.Root,
		Created_at: timeStamp,
		Father:     g.Father,
		Status:     g.Status,
		Message:    g.Message,
	}

	res := r.data.db.Table(TABLE_COMMENT).Create(&u)
	return res.Error
}

func (r *forestRepo) GetForest(ctx context.Context, g *v1.GetLeafsRequest) (list []*types.Leaf, total int64, err error) {

	var leafs []types.Leaf
	var count int64
	var res *gorm.DB

	res = r.data.db.Table(TABLE_FOREST).Order("id desc")
	res = res.Offset(int((g.Page - 1) * g.Pagesize)).Limit(int(g.Pagesize)).Find(&leafs).Count(&count)

	if res.Error != nil {
		return nil, 0, res.Error
	}

	list = make([]*types.Leaf, 0)

	for _, v := range leafs {
		if v.Status != 1 {
			v.Owner = 0
		}
		list = append(list, &types.Leaf{
			ID:        v.ID,
			Owner:     v.Owner,
			Create_at: v.Create_at,
			Status:    v.Status,
			Message:   v.Message,
		})
	}

	return list, count, nil
}
