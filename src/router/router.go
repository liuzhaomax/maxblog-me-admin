package router

import (
	"github.com/gin-gonic/gin"
	"maxblog-me-admin/src/handler"
)

func RegisterRouter(handler *handler.HUser, group *gin.RouterGroup) {
	routerUser := group.Group("")
	{
		routerUser.GET("/:id", handler.GetUserById)
	}
}
