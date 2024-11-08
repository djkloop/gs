package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Content struct {
	ID             int64         `json:"id"`
	Title          string        `json:"title" binding:"required"`
	Description    string        `json:"description" binding:"required"`
	Author         string        `json:"author" binding:"required"`
	VideoURL       string        `json:"video_url"`
	Thumbnail      string        `json:"thumbnail"`
	Category       string        `json:"category"`
	Duration       time.Duration `json:"duration"`
	Resolution     string        `json:"resolution"`
	FileSize       int64         `json:"file_size"`
	Format         string        `json:"format"`
	Quality        int32         `json:"quality"`
	ApprovalStatus int32         `json:"approval_status"`
}

type ContentRepo interface {
	Create(context.Context, *Content) (*Content, error)
	Update(context.Context, *Content) (*Content, error)
	FindByID(context.Context, int64) (*Content, error)
	ListByHello(context.Context, string) ([]*Content, error)
	ListAll(context.Context) ([]*Content, error)
}

type ContentUseCase struct {
	repo ContentRepo
	log  *log.Helper
}

func NewContentUseCase(repo ContentRepo, logger log.Logger) *ContentUseCase {
	return &ContentUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ContentUseCase) CreateContent(ctx context.Context, g *Content) (*Content, error) {
	uc.log.WithContext(ctx).Infof("CreateContent: %v", g)
	return uc.repo.Create(ctx, g)
}

func (uc *ContentUseCase) UpdateContent(ctx context.Context, g *Content) (*Content, error) {
	uc.log.WithContext(ctx).Infof("UpdateContent: %v", g)
	return uc.repo.Update(ctx, g)
}
