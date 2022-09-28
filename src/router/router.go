package router

import (
	"github.com/gin-gonic/gin"
	"maxblog-me-admin/src/handler"
)

func RegisterRouter(handler *handler.HUser, group *gin.RouterGroup) {
	routerIndex := group.Group("")
	{
		routerIndex.GET("/", handler.GetIndex)
		routerIndex.GET("/login", handler.GetPuk)
		routerIndex.POST("/login", handler.PostLogin)
	}
	routerUser := group.Group("/user")
	{
		routerUser.GET("/:id", handler.GetUserById)
	}
}
