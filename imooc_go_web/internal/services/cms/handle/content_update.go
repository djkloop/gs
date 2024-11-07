package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"imooc_go_web/internal/dao"
	"imooc_go_web/internal/model"
	"net/http"
	"time"
)

type ContentUpdateReq struct {
	ID             int           `json:"id" binding:"required"`
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

type ContentUpdateResp struct {
	Message string `json:"message"`
}

func (h *CmsHandle) ContentUpdateHandle(c *gin.Context) {
	var req ContentUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	newContentDao := dao.NewContentDao(h.DB)

	// 判断文章是否存在
	ok, err := newContentDao.IsExist(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "content not exist",
		})
		return
	}

	if err := newContentDao.Update(req.ID, &model.ContentDetail{
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
	}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": &ContentUpdateResp{
			Message: "ok",
		},
	})
}
