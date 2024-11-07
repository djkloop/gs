package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"imooc_go_web/internal/dao"
	"imooc_go_web/internal/model"
	"net/http"
	"time"
)

type RegisterReq struct {
	UserId   string `json:"user_id" binding:"required"`
	NickName string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRespData struct {
	Msg string `json:"message"`
}

func (h *CmsHandle) RegisterHandle(c *gin.Context) {

	var req RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	// 加密密码
	hashPassword, err := encryptPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// db 操作
	newAccountDao := dao.NewAccountDao(h.DB)

	// 检查账号是否存在
	isExist, err := newAccountDao.IsExist(req.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "register fail",
		})
		return
	}
	if isExist {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data": &RegisterRespData{
				Msg: fmt.Sprintf("账号已存在"),
			},
		})
		return
	}

	// 账号存入数据库
	nowTime := time.Now()
	err = newAccountDao.Create(&model.Account{
		UserId:    req.UserId,
		Nickname:  req.NickName,
		Password:  hashPassword,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "register fail",
		})
		return
	}

	fmt.Printf("register req = %+v , hashPassword = [%s] \n", req, hashPassword)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": &RegisterRespData{
			Msg: fmt.Sprintf("注册成功"),
		},
	})
}

// 加密密码
func encryptPassword(password string) (string, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Printf("密码加密失败, err = %+v \n", err)
		return "", err
	}

	return string(hashPassword), nil
}
