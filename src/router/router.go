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
	routerMaxBlog := app.Group("/maxblog")
	{
		routerMaxBlog.Use(itcpt.InterceptorAuth.CheckTokens())

		routerHome := routerMaxBlog.Group("/home")
		{
			routerHome.GET("", handler.GetHome)
		}

		routerUser := routerMaxBlog.Group("/user")
		{
			routerUser.GET("/:id", handler.GetUserById)
			routerUser.POST("", handler.PostUser)
		}
	}
}
