package data

import (
	"content_manage/internal/biz"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}

type ContentDetail struct {
	ID             int64         `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 内容ID
	Title          string        `gorm:"column:title"`                         // 内容标题
	Description    string        `gorm:"column:description"`                   // 内容描述
	Author         string        `gorm:"column:author"`                        // 作者
	VideoURL       string        `gorm:"column:video_url"`                     // 视频URL
	Thumbnail      string        `gorm:"column:thumbnail"`                     // 缩略图URL
	Category       string        `gorm:"column:category"`                      // 分类
	Duration       time.Duration `gorm:"column:duration"`                      // 时长
	Resolution     string        `gorm:"column:resolution"`                    // 分辨率
	FileSize       int64         `gorm:"column:file_size"`                     // 文件大小
	Format         string        `gorm:"column:format"`                        // 文件格式 1 - MP4  2 - FLV  3 - AVI
	Quality        int32         `gorm:"column:quality"`                       // 视频质量 1 - 高清  2 - 标清
	ApprovalStatus int32         `gorm:"column:approval_status"`               // 审批状态 1 - 审核中  2 - 审核通过 3 - 审核失败
	CreatedAt      time.Time     `gorm:"column:created_at"`                    // 创建时间
	UpdatedAt      time.Time     `gorm:"column:updated_at"`                    // 更新时间
}

func (d ContentDetail) TableName() string {
	return "cms_content_detail"
}

func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *contentRepo) Create(ctx context.Context, g *biz.Content) (*biz.Content, error) {
	db := r.data.DB
	saveContent := &ContentDetail{
		Title:          g.Title,
		Description:    g.Description,
		Author:         g.Author,
		VideoURL:       g.VideoURL,
		Thumbnail:      g.Thumbnail,
		Category:       g.Category,
		Duration:       g.Duration,
		Resolution:     g.Resolution,
		FileSize:       g.FileSize,
		Format:         g.Format,
		Quality:        g.Quality,
		ApprovalStatus: g.ApprovalStatus,
	}

	if err := db.Create(&saveContent).Error; err != nil {
		r.log.Errorf("CreateContent: %v", err)
		return nil, err
	}

	return g, nil
}

func (r *contentRepo) Update(ctx context.Context, g *biz.Content) (*biz.Content, error) {
	db := r.data.DB
	saveContent := &ContentDetail{
		ID:             g.ID,
		Title:          g.Title,
		Description:    g.Description,
		Author:         g.Author,
		VideoURL:       g.VideoURL,
		Thumbnail:      g.Thumbnail,
		Category:       g.Category,
		Duration:       g.Duration,
		Resolution:     g.Resolution,
		FileSize:       g.FileSize,
		Format:         g.Format,
		Quality:        g.Quality,
		ApprovalStatus: g.ApprovalStatus,
	}

	if err := db.Updates(&saveContent).Error; err != nil {
		r.log.Errorf("UpdateContent: %v", err)
		return nil, err
	}
	return g, nil
}

func (r *contentRepo) FindByID(context.Context, int64) (*biz.Content, error) {
	return nil, nil
}

func (r *contentRepo) ListByHello(context.Context, string) ([]*biz.Content, error) {
	return nil, nil
}

func (r *contentRepo) ListAll(context.Context) ([]*biz.Content, error) {
	return nil, nil
}
