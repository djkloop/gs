package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"imooc_go_web/internal/api"
)

func main() {
	r := gin.Default()
	api.CmsRouters(r)
	err := r.Run()
	if err != nil {
		fmt.Println("", err)
		return
	}
}
