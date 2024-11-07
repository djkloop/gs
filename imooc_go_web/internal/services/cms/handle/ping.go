package handle

import "github.com/gin-gonic/gin"

func (h *CmsHandle) PingHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
