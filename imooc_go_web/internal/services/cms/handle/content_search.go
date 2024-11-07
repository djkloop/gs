package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"imooc_go_web/internal/dao"
	"imooc_go_web/internal/model"
	"net/http"
)

type ContentSearchReq struct {
	ID       int    `json:"id"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type ContentSearchResp struct {
	Message string                 `json:"message"`
	List    []*model.ContentDetail `json:"list"`
	Total   int64                  `json:"total"`
}

func (h *CmsHandle) ContentSearchHandle(c *gin.Context) {
	var req ContentSearchReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	newContentDao := dao.NewContentDao(h.DB)

	// 判断文章是否存在
	contents, total, err := newContentDao.Search(&dao.SearchParams{
		ID:       req.ID,
		Author:   req.Author,
		Title:    req.Title,
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": &ContentSearchResp{
			Message: "ok",
			List:    contents,
			Total:   total,
		},
	})
}
