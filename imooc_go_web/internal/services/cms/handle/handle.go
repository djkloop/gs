package handle

import (
	"github.com/redis/go-redis/v9"
	goflow "github.com/s8sg/goflow/v1"
	"gorm.io/gorm"
)

type CmsHandle struct {
	DB          *gorm.DB
	RDB         *redis.Client
	FlowService *goflow.FlowService
}

func NewCmsHandle(db *gorm.DB, rdb *redis.Client, flowService *goflow.FlowService) *CmsHandle {
	return &CmsHandle{
		DB:          db,
		RDB:         rdb,
		FlowService: flowService,
	}
}

//func (h *CmsHandle) LoginHandle(c *gin.Context) {
//	LoginHandle(c)
//}

//func (h *CmsHandle) PingHandle(c *gin.Context) {
//	PingHandle(c)
//}
