package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"imooc_go_web/internal/dao"
	"net/http"
)

type ContentDeleteReq struct {
	ID int `json:"id" binding:"required"`
}

type ContentDeleteResp struct {
	Message string `json:"message"`
}

func (h *CmsHandle) ContentDeleteHandle(c *gin.Context) {
	var req ContentDeleteReq
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

	if err := newContentDao.Delete(req.ID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": &ContentDeleteResp{
			Message: "ok",
		},
	})
}
