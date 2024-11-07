package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"imooc_go_web/internal/utils"
	"net/http"
)

const (
	SessionKey = "session_id"
)

type SessionMiddleware struct {
	RDB *redis.Client
}

func NewSessionAuthMiddleware() *SessionMiddleware {
	s := &SessionMiddleware{}
	connectMiddleWareRDB(s)
	return s
}

func (s *SessionMiddleware) Auth(c *gin.Context) {
	sessionId := c.GetHeader(SessionKey)
	ctx := context.Background()
	// TODO 校验sessionId是否合法
	if sessionId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "sessionId is empty",
		})
		return
	}

	authKey := utils.GetAuthKey(sessionId)
	loginTime, err := s.RDB.Get(ctx, authKey).Result()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "登录失败!",
		})
		return
	}
	if loginTime == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "账户过期!",
		})
		return
	}

	fmt.Println("sessionId:", sessionId)
	c.Next()
}

func connectMiddleWareRDB(middleware *SessionMiddleware) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "43.143.243.166:6379",
		Password: "bS9@xG2?", // no password set
		DB:       0,          // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	middleware.RDB = rdb
}
