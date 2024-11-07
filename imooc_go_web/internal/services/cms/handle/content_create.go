package handle

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	goflow "github.com/s8sg/goflow/v1"
	"imooc_go_web/internal/dao"
	"imooc_go_web/internal/model"
	"net/http"
	"time"
)

type ContentCreateReq struct {
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
	Quality        int           `json:"quality"`
	ApprovalStatus int           `json:"approval_status"`
}

type ContentCreateResp struct {
	Message string `json:"message"`
}

func (h *CmsHandle) ContentCreateHandle(c *gin.Context) {
	var req ContentCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	newContentDao := dao.NewContentDao(h.DB)
	id, err := newContentDao.Create(&model.ContentDetail{
		Title:          req.Title,
		Description:    req.Description,
		Author:         req.Description,
		VideoURL:       req.VideoURL,
		Thumbnail:      req.Thumbnail,
		Category:       req.Category,
		Duration:       req.Duration,
		Resolution:     req.Resolution,
		FileSize:       req.FileSize,
		Format:         req.Format,
		Quality:        req.Quality,
		ApprovalStatus: req.ApprovalStatus,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	// 工作流
	flowServiceData := map[string]interface{}{
		"content_id": id,
	}
	data, _ := json.Marshal(flowServiceData)
	if err := h.FlowService.Execute("content-flow", &goflow.Request{
		Body: data,
	}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": &ContentCreateResp{
			Message: "ok",
		},
	})
}
