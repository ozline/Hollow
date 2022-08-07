package biz

import (
	"context"

	v1 "hollow/api/hollow/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type ForestRepo interface {
	PushLeaf(ctx context.Context, g *v1.PushLeafRequest) error
	GetForest(ctx context.Context, g *v1.GetLeafsRequest) (list []*Leaf, total int64, err error)
}

type ForestUsecase struct {
	ur  ForestRepo
	log *log.Helper
}

type Leaf struct {
	ID        int64
	Owner     int64
	Create_at int64
	Status    int64
	Message   string
}

func converttoReply(forest *Leaf) *v1.Leaf {
	return &v1.Leaf{
		Id:       forest.ID,
		Owner:    forest.Owner,
		CreateAt: forest.Create_at,
		Status:   forest.Status,
		Message:  forest.Message,
	}
}

func NewForestUsecase(repo ForestRepo, logger log.Logger) *ForestUsecase {
	return &ForestUsecase{ur: repo, log: log.NewHelper(logger)}
}

func (uc *ForestUsecase) PushLeaf(ctx context.Context, u *v1.PushLeafRequest) error {

	return uc.ur.PushLeaf(ctx, u)
}

func (uc *ForestUsecase) GetForest(ctx context.Context, u *v1.GetLeafsRequest) (list []*v1.Leaf, total int64, err error) {

	leafList, count, err := uc.ur.GetForest(ctx, u)

	if err != nil {
		return nil, 0, err
	}

	forest := make([]*v1.Leaf, 0)

	for _, v := range leafList {
		forest = append(forest, converttoReply(v))
	}

	return forest, count, nil

}
