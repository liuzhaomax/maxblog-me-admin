package router

import (
	"github.com/gin-gonic/gin"
	"maxblog-me-admin/internal/middleware/interceptor"
	"maxblog-me-admin/src/handler"
)

func RegisterRouter(handler *handler.HUser, group *gin.RouterGroup) {
	routerIndex := group.Group("")
	{
		routerIndex.GET("/", handler.GetIndex)
		routerIndex.GET("/login", handler.GetPuk)
		routerIndex.POST("/login", handler.PostLogin)
		routerIndex.DELETE("/logout", handler.DeleteLogout)
	}

	itcpt := interceptor.GetInstanceOfContext()

	routerUser := group.Group("/user")
	{
		routerUser.Use(itcpt.CheckTokens())

		routerUser.GET("/:id", handler.GetUserById)
		routerUser.POST("", handler.PostUser)
	}
}
