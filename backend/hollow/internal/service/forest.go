package service

import (
	"context"
	v1 "hollow/api/hollow/v1"
	"hollow/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	errors "hollow/internal/errors"
)

type ForestService struct {
	v1.UnimplementedForestsServer
	uc  *biz.ForestUsecase
	log *log.Helper
}

func NewForestService(uc *biz.ForestUsecase) *ForestService {
	return &ForestService{uc: uc}
}

// 推送叶子
func (s *ForestService) PushLeaf(ctx context.Context, req *v1.PushLeafRequest) (reply *v1.PushLeafReply, err error) {
	if len(req.Message) > 140 || len(req.Message) == 0 {
		return nil, errors.ErrParamsIllegal
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

// 删除叶子
func (s *ForestService) DeleteLeaf(ctx context.Context, req *v1.DeleteLeafRequest) (reply *v1.DeleteLeafReply, err error) {

	err = s.uc.DeleteLeaf(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.DeleteLeafReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

// 获取叶子
func (s *ForestService) GetForest(ctx context.Context, req *v1.GetLeafsRequest) (reply *v1.GetLeafsReply, err error) {
	if req.Page == 0 || req.Pagesize == 0 {
		return nil, errors.ErrParamsIllegal
	}

	forest, count, err := s.uc.GetForest(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.GetLeafsReply{
		Code: 200,
		Msg:  "ok",
		Data: &v1.MultipleLeafReply{
			List:  forest,
			Total: count,
		},
	}, nil
}

// 获取叶子信息
func (s *ForestService) GetLeafDetail(ctx context.Context, req *v1.GetLeafDetailRequest) (reply *v1.GetLeafDetailReply, err error) {

	leaf, err := s.uc.GetLeafDetail(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.GetLeafDetailReply{
		Code: 200,
		Msg:  "ok",
		Data: &v1.Leaf{
			Id:        leaf.ID,
			Owner:     leaf.Owner,
			Message:   leaf.Message,
			CreatedAt: leaf.Created_at,
			Status:    leaf.Status,
			Liked:     leaf.Liked,
		},
	}, nil
}

// 评论
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

// 获取评论
func (s *ForestService) GetComments(ctx context.Context, req *v1.GetCommentsRequest) (reply *v1.GetCommentsReply, err error) {

	comments, count, err := s.uc.GetComments(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.GetCommentsReply{
		Code: 200,
		Msg:  "ok",
		Data: &v1.MultipleCommentReply{
			List:  comments,
			Total: count,
		},
	}, nil
}

// 删除评论
func (s *ForestService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (reply *v1.DeleteCommentReply, err error) {
	err = s.uc.DeleteComment(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.DeleteCommentReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

// 点赞评论
func (s *ForestService) LikeComment(ctx context.Context, req *v1.LikeCommentRequest) (reply *v1.LikeCommentReply, err error) {
	err = s.uc.LikeComment(ctx, req)

	if err != nil {
		return nil, err
	}

	return &v1.LikeCommentReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}
