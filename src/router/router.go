package router

import (
	"github.com/gin-gonic/gin"
	"maxblog-me-admin/internal/middleware/interceptor"
	"maxblog-me-admin/src/handler"
)

func RegisterRouter(handler *handler.HUser, app *gin.Engine, itcpt *interceptor.Interceptor) {
	routerIndex := app.Group("")
	{
		routerIndex.GET("/", handler.GetIndex)
		routerIndex.GET("/login", handler.GetPuk)
		routerIndex.POST("/login", handler.PostLogin)
		routerIndex.DELETE("/logout", handler.DeleteLogout)
	}

	routerHome := app.Group("")
	{
		routerHome.Use(itcpt.InterceptorAuth.CheckTokens())
		routerHome.GET("/home", handler.GetHome)
	}

	routerUser := app.Group("")
	{
		routerUser.Use(itcpt.InterceptorAuth.CheckTokens())
		routerUser.GET("/user/:id", handler.GetUserById)
		routerUser.POST("/user", handler.PostUser)
	}
}
