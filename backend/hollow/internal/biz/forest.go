package biz

import (
	"context"

	v1 "hollow/api/hollow/v1"

	"hollow/internal/errors"
	models "hollow/internal/models"

	"github.com/go-kratos/kratos/v2/log"
)

type ForestRepo interface {
	PushLeaf(ctx context.Context, g *v1.PushLeafRequest) error
	DeleteLeaf(ctx context.Context, g *v1.DeleteLeafRequest) error
	GetForest(ctx context.Context, g *v1.GetLeafsRequest) (list []*models.Leaf, total int64, err error)
	GetLeafDetail(ctx context.Context, g *v1.GetLeafDetailRequest) (leaf *models.Leaf, err error)
	CommentLeaf(ctx context.Context, g *v1.CommentLeafRequest) error
	GetComments(ctx context.Context, g *v1.GetCommentsRequest) (list []*models.Comment, total int64, err error)
	DeleteComment(ctx context.Context, g *v1.DeleteCommentRequest) error
	LikeComment(ctx context.Context, u *v1.LikeCommentRequest) error
	Report(ctx context.Context, g *v1.ReportRequest) error
	GetReportList(ctx context.Context, g *v1.GetReportListRequest) (list []*models.Report, total int64, err error)
	UpdateReport(ctx context.Context, g *v1.UpdateReportRequest) error
	UpdateCommentStatus(ctx context.Context, g *v1.UpdateCommentStatusRequest) error
	UpdateLeafStatus(ctx context.Context, g *v1.UpdateLeafStatusRequest) error
}

type ForestUsecase struct {
	ur  ForestRepo
	log *log.Helper
}

func convertLeafToReply(forest *models.Leaf) *v1.Leaf {
	return &v1.Leaf{
		Id:        forest.ID,
		Owner:     forest.Owner,
		CreatedAt: forest.Created_at,
		Status:    forest.Status,
		Message:   forest.Message,
		Liked:     forest.Liked,
	}
}

func convertCommentToReply(comment *models.Comment) *v1.Comment {
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

func convertReportReply(report *models.Report) *v1.Report {
	return &v1.Report{
		Id:        report.Id,
		Type:      report.Type,
		Status:    report.Status,
		Reporter:  report.Reporter,
		ReportId:  report.Report_id,
		Reason:    report.Reason,
		Message:   report.Message,
		Reply:     report.Reply,
		CreatedAt: report.Created_at,
		UpdatedAt: report.Updated_at,
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

func (uc *ForestUsecase) GetLeafDetail(ctx context.Context, u *v1.GetLeafDetailRequest) (comment *models.Leaf, err error) {
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

func (uc *ForestUsecase) Report(ctx context.Context, u *v1.ReportRequest) error {
	return uc.ur.Report(ctx, u)
}

func (uc *ForestUsecase) GetReportList(ctx context.Context, u *v1.GetReportListRequest) (list []*v1.Report, total int64, err error) {

	reportList, count, err := uc.ur.GetReportList(ctx, u)

	if err != nil {
		return nil, 0, errors.GenerateErrorNormal(err)
	}

	reports := make([]*v1.Report, 0)

	for _, v := range reportList {
		reports = append(reports, convertReportReply(v))
	}

	return reports, count, nil
}

func (uc *ForestUsecase) UpdateReport(ctx context.Context, u *v1.UpdateReportRequest) error {
	return uc.ur.UpdateReport(ctx, u)
}

func (uc *ForestUsecase) UpdateCommentStatus(ctx context.Context, u *v1.UpdateCommentStatusRequest) error {
	return uc.ur.UpdateCommentStatus(ctx, u)
}

func (uc *ForestUsecase) UpdateLeafStatus(ctx context.Context, u *v1.UpdateLeafStatusRequest) error {
	return uc.ur.UpdateLeafStatus(ctx, u)
}
