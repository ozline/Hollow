package data

import (
	"context"
	"fmt"

	v1 "hollow/api/hollow/v1"
	"hollow/internal/biz"

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
	timeStamp := biz.GetTimestamp13()
	// userid := getUserid(ctx)
	userid := int64(26)

	u := biz.Leaf{
		Owner:     userid,
		Create_at: timeStamp,
		Status:    g.Status,
		Message:   g.Message,
	}

	fmt.Println(u)

	res := r.data.db.Table(TABLE_FOREST).Create(&u)

	return res.Error
}

func (r *forestRepo) GetForest(ctx context.Context, g *v1.GetLeafsRequest) (list []*biz.Leaf, total int64, err error) {

	var leafs []biz.Leaf
	var count int64
	var res *gorm.DB

	fmt.Println("用户id:", getUserid(ctx))

	res = r.data.db.Table(TABLE_FOREST).Order("id desc")
	res = res.Offset(int((g.Page - 1) * g.Pagesize)).Limit(int(g.Pagesize)).Find(&leafs).Count(&count)

	if res.Error != nil {
		return nil, 0, res.Error
	}

	list = make([]*biz.Leaf, 0)

	for _, v := range leafs {
		list = append(list, &biz.Leaf{
			ID:        v.ID,
			Owner:     v.Owner,
			Create_at: v.Create_at,
			Status:    v.Status,
			Message:   v.Message,
		})
	}

	return list, count, nil
}
