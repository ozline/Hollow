package service

import (
	"context"
	v1 "hollow/api/hollow/v1"
	"hollow/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ForestService struct {
	v1.UnimplementedForestsServer
	uc  *biz.ForestUsecase
	log *log.Helper
}

func NewForestService(uc *biz.ForestUsecase) *ForestService {
	return &ForestService{uc: uc}
}

//推送叶子
func (s *ForestService) Push(ctx context.Context, req *v1.PushLeafRequest) (reply *v1.PushLeafReply, err error) {
	if len(req.Message) > 140 || len(req.Message) == 0 {
		return nil, ErrParamsIllegal
	}

	err = s.uc.PushLeaf(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.PushLeafReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

//获取叶子
func (s *ForestService) Get(ctx context.Context, req *v1.GetLeafsRequest) (reply *v1.GetLeafsReply, err error) {
	if req.Page == 0 || req.Pagesize == 0 {
		return nil, ErrParamsIllegal
	}

	forest, count, err := s.uc.GetForest(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.GetLeafsReply{
		Code: 200,
		Msg:  "ok",
		Data: &v1.MultipleTodoReply{
			List:  forest,
			Total: count,
		},
	}, nil
}

//评论
func (s *ForestService) Comment(ctx context.Context, req *v1.CommentLeafRequest) (reply *v1.CommentLeafRePly, err error) {

	err = s.uc.CommentLeaf(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.CommentLeafRePly{
		Code: 200,
		Msg:  "ok",
	}, nil
}
