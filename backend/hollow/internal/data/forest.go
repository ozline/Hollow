package data

import (
	"context"

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

// 查询某一个东西是否存在
func (r *forestRepo) isExist(table string, query interface{}, args ...interface{}) bool {
	var count int64 = 0
	_ = r.data.db.Table(table).Where(query, args).Limit(1).Count(&count)
	return count != 0
}

// 发表树叶
func (r *forestRepo) PushLeaf(ctx context.Context, g *v1.PushLeafRequest) error {
	timeStamp := utils.GetTimestamp13()
	user := GetUserInfo(ctx) //获取用户信息

	if user.ID == -1 {
		return errors.ErrUserInvalid
	}

	u := types.Leaf{
		ID:         GetSnowflakeID(r.data.node),
		Owner:      user.ID,
		Created_at: timeStamp,
		Status:     g.Status,
		Message:    g.Message,
	}
	res := r.data.db.Table(TABLE_FOREST).Create(&u)
	return res.Error
}

// 删除评论
func (r *forestRepo) DeleteLeaf(ctx context.Context, g *v1.DeleteLeafRequest) error {
	leaf := new(types.Leaf)
	res := r.data.db.Table(TABLE_FOREST).Where("id = ?", g.Id).Limit(1).Find(&leaf)

	if res.Error != nil {
		return res.Error
	}

	user := GetUserInfo(ctx)
	if user.ID == -1 {
		return errors.ErrUserInvalid
	}

	if user.ID != leaf.Owner {
		return errors.ErrUserNotMatched
	}

	res = res.Unscoped().Delete(&leaf)

	return res.Error
}

// 获取列表
func (r *forestRepo) GetForest(ctx context.Context, g *v1.GetLeafsRequest) (list []*types.Leaf, total int64, err error) {
	var leafs []types.Leaf
	var count int64
	var res *gorm.DB

	res = r.data.db.Table(TABLE_FOREST).Order("id desc").Count(&count)
	res = res.Offset(int((g.Page - 1) * g.Pagesize)).Limit(int(g.Pagesize)).Find(&leafs)

	if res.Error != nil {
		return nil, 0, res.Error
	}

	list = make([]*types.Leaf, 0)

	for _, v := range leafs {
		if v.Status != 1 { //匿名 且 与自身ID不等
			v.Owner = 0
		}
		list = append(list, &types.Leaf{
			ID:         v.ID,
			Owner:      v.Owner,
			Created_at: v.Created_at,
			Status:     v.Status,
			Message:    v.Message,
		})
	}

	return list, count, nil
}

func (r *forestRepo) GetLeafDetail(ctx context.Context, g *v1.GetLeafDetailRequest) (leaf *types.Leaf, err error) {
	var data = new(types.Leaf)
	res := r.data.db.Table(TABLE_FOREST).Where("id = ?", g.Id).Limit(1).Find(&data)

	if res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

// 跟帖
func (r *forestRepo) CommentLeaf(ctx context.Context, g *v1.CommentLeafRequest) error {
	timeStamp := utils.GetTimestamp13()
	user := GetUserInfo(ctx)

	if user.ID == -1 {
		return errors.ErrUserInvalid
	}

	// 验证父节点是否存在
	if g.Father != 0 {
		if !r.isExist(TABLE_USERS, "id = ?", g.Father) {
			return errors.ErrFatherNotExisted
		}
	}

	// 验证楼层是否存在
	if !r.isExist(TABLE_FOREST, "id = ?", g.Root) {
		return errors.ErrLeafNotExisted
	}

	u := types.Comment{
		ID:         GetSnowflakeID(r.data.node),
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

func (r *forestRepo) GetComments(ctx context.Context, g *v1.GetCommentsRequest) (list []*types.Comment, total int64, err error) {
	var comments []types.Comment
	var count int64
	var res *gorm.DB

	res = r.data.db.Table(TABLE_COMMENT).Where("root = ? AND father = ?", g.Root, g.Father).Count(&count)
	res = res.Offset(int((g.Page - 1) * g.Pagesize)).Limit(int(g.Pagesize)).Find(&comments)

	if res.Error != nil {
		return nil, 0, res.Error
	}

	user := GetUserInfo(ctx)

	list = make([]*types.Comment, 0)

	for _, v := range comments {
		if v.Status != 1 && v.Owner != user.ID { //匿名 且 与自身ID不等
			v.Owner = 0
		}
		list = append(list, &types.Comment{
			ID:         v.ID,
			Owner:      v.Owner,
			Root:       v.Root,
			Father:     v.Father,
			Created_at: v.Created_at,
			Status:     v.Status,
			Message:    v.Message,
			Liked:      v.Liked,
		})
	}

	return list, count, nil
}

func (r *forestRepo) DeleteComment(ctx context.Context, g *v1.DeleteCommentRequest) error {
	comment := new(types.Comment)
	var count int64
	res := r.data.db.Table(TABLE_COMMENT).Where("id = ?", g.Id).Limit(1).Count(&count)

	if res.Error != nil {
		return res.Error
	}

	if count == 0 {
		return errors.ErrCommentNotFound
	}

	res = res.First(&comment)

	user := GetUserInfo(ctx)
	if user.ID == -1 {
		return errors.ErrUserInvalid
	}

	if user.ID != comment.Owner {
		return errors.ErrUserNotMatched
	}

	// res = res.Unscoped().Delete(&comment)

	comment.Status = 2 //标记为删除

	res = res.Save(&comment)

	return res.Error
}

func (r *forestRepo) LikeComment(ctx context.Context, g *v1.LikeCommentRequest) error {
	var count int64
	res := r.data.db.Table(TABLE_COMMENT).Where("id = ?", g.Id).Limit(1).Count(&count)

	if res.Error != nil {
		return res.Error
	}

	if count == 0 {
		return errors.ErrCommentNotFound
	}

	user := GetUserInfo(ctx)

	if user.ID == -1 {
		return errors.ErrUserInvalid
	}

	res = r.data.db.Table(TABLE_LIKE).Where("comment = ? AND user = ?", g.Id, user.ID).Limit(1).Count(&count)

	if count != 0 {
		return errors.ErrHaveLiked
	}

	like := types.Like{
		Timestamp: utils.GetTimestamp13(),
		User:      user.ID,
		Comment:   g.Id,
	}

	res = res.Create(&like)

	if res.Error != nil {
		return errors.GenerateErrorNormal(res.Error)
	}

	comment := new(types.Comment)

	res = r.data.db.Table(TABLE_COMMENT).Where("id = ?", g.Id).First(&comment)

	comment.Liked = comment.Liked + 1

	res = res.Save(&comment)

	return res.Error
}
