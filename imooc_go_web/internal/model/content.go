package model

import "time"

type ContentDetail struct {
	ID             int           `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 内容ID
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
	Quality        int           `gorm:"column:quality"`                       // 视频质量 1 - 高清  2 - 标清
	ApprovalStatus int           `gorm:"column:approval_status"`               // 审批状态 1 - 审核中  2 - 审核通过 3 - 审核失败
	CreatedAt      time.Time     `gorm:"column:created_at"`                    // 创建时间
	UpdatedAt      time.Time     `gorm:"column:updated_at"`                    // 更新时间
}

func (d ContentDetail) TableName() string {
	return "cms_content_detail"
}
