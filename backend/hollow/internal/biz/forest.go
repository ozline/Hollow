package biz

import (
	"context"

	v1 "hollow/api/hollow/v1"

	"hollow/internal/errors"
	types "hollow/internal/types"

	"github.com/go-kratos/kratos/v2/log"
)

type ForestRepo interface {
	PushLeaf(ctx context.Context, g *v1.PushLeafRequest) error
	DeleteLeaf(ctx context.Context, g *v1.DeleteLeafRequest) error
	GetForest(ctx context.Context, g *v1.GetLeafsRequest) (list []*types.Leaf, total int64, err error)
	GetLeafDetail(ctx context.Context, g *v1.GetLeafDetailRequest) (leaf *types.Leaf, err error)
	CommentLeaf(ctx context.Context, g *v1.CommentLeafRequest) error
	GetComments(ctx context.Context, g *v1.GetCommentsRequest) (list []*types.Comment, total int64, err error)
	DeleteComment(ctx context.Context, g *v1.DeleteCommentRequest) error
	LikeComment(ctx context.Context, u *v1.LikeCommentRequest) error
}

type ForestUsecase struct {
	ur  ForestRepo
	log *log.Helper
}

func convertLeafToReply(forest *types.Leaf) *v1.Leaf {
	return &v1.Leaf{
		Id:        forest.ID,
		Owner:     forest.Owner,
		CreatedAt: forest.Created_at,
		Status:    forest.Status,
		Message:   forest.Message,
		Liked:     forest.Liked,
	}
}

func convertCommentToReply(comment *types.Comment) *v1.Comment {
	return &v1.Comment{
		Id:        comment.ID,
		Owner:     comment.Owner,
		Father:    comment.Father,
		Root:      comment.Root,
		CreatedAt: comment.Created_at,
		Status:    comment.Status,
		Message:   comment.Message,
		Liked:     comment.Liked,
	}
}

func NewForestUsecase(repo ForestRepo, logger log.Logger) *ForestUsecase {
	return &ForestUsecase{ur: repo, log: log.NewHelper(logger)}
}

func (uc *ForestUsecase) PushLeaf(ctx context.Context, u *v1.PushLeafRequest) error {
	return uc.ur.PushLeaf(ctx, u)
}

func (uc *ForestUsecase) DeleteLeaf(ctx context.Context, u *v1.DeleteLeafRequest) error {
	return uc.ur.DeleteLeaf(ctx, u)
}

func (uc *ForestUsecase) GetForest(ctx context.Context, u *v1.GetLeafsRequest) (list []*v1.Leaf, total int64, err error) {

	leafList, count, err := uc.ur.GetForest(ctx, u)

	if err != nil {
		return nil, 0, err
	}

	forest := make([]*v1.Leaf, 0)

	for _, v := range leafList {
		forest = append(forest, convertLeafToReply(v))
	}

	return forest, count, nil
}

func (uc *ForestUsecase) GetLeafDetail(ctx context.Context, u *v1.GetLeafDetailRequest) (comment *types.Leaf, err error) {
	return uc.ur.GetLeafDetail(ctx, u)
}

func (uc *ForestUsecase) CommentLeaf(ctx context.Context, u *v1.CommentLeafRequest) error {
	return uc.ur.CommentLeaf(ctx, u)
}

func (uc *ForestUsecase) GetComments(ctx context.Context, u *v1.GetCommentsRequest) (list []*v1.Comment, total int64, err error) {
	commentsList, count, err := uc.ur.GetComments(ctx, u)

	if err != nil {
		return nil, 0, errors.GenerateErrorNormal(err)
	}

	comments := make([]*v1.Comment, 0)

	for _, v := range commentsList {
		comments = append(comments, convertCommentToReply(v))
	}

	return comments, count, nil
}

func (uc *ForestUsecase) DeleteComment(ctx context.Context, u *v1.DeleteCommentRequest) error {
	return uc.ur.DeleteComment(ctx, u)
}

func (uc *ForestUsecase) LikeComment(ctx context.Context, u *v1.LikeCommentRequest) error {
	return uc.ur.LikeComment(ctx, u)
}
