package api

import (
	"github.com/gin-gonic/gin"
	"imooc_go_web/internal/services"
)

const (
	apiRootPath      = "/api"
	apiV1Path        = "/v1"
	apiV1OutPath     = "/out"
	apiRootV1CmsPath = apiRootPath + apiV1Path + "/cms"
	apiOutV1CmsPath  = apiRootPath + apiV1Path + apiV1OutPath + "/cms"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := services.NewCmsApp()
	//
	sessionMiddleware := NewSessionAuthMiddleware()
	apiRootV1CmsGroup := r.Group(apiRootV1CmsPath).Use(sessionMiddleware.Auth)
	{
		apiRootV1CmsGroup.GET("/ping", cmsApp.CmsHandle.PingHandle)
		apiRootV1CmsGroup.POST("/content/create", cmsApp.CmsHandle.ContentCreateHandle)
		apiRootV1CmsGroup.POST("/content/update", cmsApp.CmsHandle.ContentUpdateHandle)
		apiRootV1CmsGroup.POST("/content/delete", cmsApp.CmsHandle.ContentDeleteHandle)
		apiRootV1CmsGroup.POST("/content/search", cmsApp.CmsHandle.ContentSearchHandle)
	}

	apiOutV1CmsPathGroup := r.Group(apiOutV1CmsPath)
	{
		apiOutV1CmsPathGroup.POST("/register", cmsApp.CmsHandle.RegisterHandle)
		apiOutV1CmsPathGroup.POST("/login", cmsApp.CmsHandle.LoginHandle)
	}
}
