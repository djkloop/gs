package handle

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"imooc_go_web/internal/dao"
	"imooc_go_web/internal/utils"
	"net/http"
	"time"
)

type LoginReq struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRespData struct {
	SessionID string `json:"session_id"`
	UserId    string `json:"user_id"`
	NickName  string `json:"nick_name"`
}

func (h *CmsHandle) LoginHandle(c *gin.Context) {

	var loginReq LoginReq
	ctx := context.Background()

	if err := c.ShouldBind(&loginReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	// 初始化dao
	newAccountDao := dao.NewAccountDao(h.DB)
	account, err := newAccountDao.FirstByUserID(loginReq.UserId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    fmt.Sprintf("账号不存在"),
		})
		return
	}

	// 密码校验
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(loginReq.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    fmt.Sprintf("密码错误"),
		})
		return
	}

	// 获取session id
	sessionId, err := h.generateSessionID(ctx, account.UserId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    fmt.Sprintf("系统错误，请稍后重试"),
		})
		return
	}

	// 响应
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": &LoginRespData{
			NickName:  account.Nickname,
			UserId:    account.UserId,
			SessionID: sessionId,
		},
	})
}

func (h *CmsHandle) generateSessionID(ctx context.Context, userID string) (string, error) {
	sessionID := uuid.New().String()
	// 生成 session id
	sessionKey := utils.GetSessionKey(userID)
	// session id 持久化
	if err := h.RDB.Set(ctx, sessionKey, sessionID, time.Hour*8).Err(); err != nil {
		fmt.Printf("redis set err error = %v \n", err)
		return "", err
	}
	// session id 过期时间
	authKey := utils.GetAuthKey(sessionID)
	if err := h.RDB.Set(ctx, authKey, time.Now().Unix(), time.Hour*8).Err(); err != nil {
		fmt.Printf("redis expire err error = %v \n", err)
		return "", err
	}

	fmt.Printf("sessionID = %s \n", sessionID)
	fmt.Printf("authKey = %s \n", authKey)

	return sessionID, nil
}
