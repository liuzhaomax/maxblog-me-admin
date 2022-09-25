package router

import (
	"github.com/gin-gonic/gin"
	"maxblog-me-admin/src/handler"
)

func RegisterRouter(handler *handler.HData, group *gin.RouterGroup) {
	routerData := group.Group("")
	{
		routerData.GET("/:id", handler.GetDataById)
	}
}
