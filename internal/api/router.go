package api

import (
	"github.com/gin-gonic/gin"
	mw "maxblog-me-template/internal/middleware"
	dataRouter "maxblog-me-template/src/router"
	"net/http"
)

func (handler *Handler) RegisterRouter(app *gin.Engine) {
	app.NoRoute(handler.GetNoRoute)
	app.Use(mw.Cors())

	router := app.Group("")
	{
		router.StaticFS("/static", http.Dir("./static"))
		dataRouter.RegisterRouter(handler.HandlerData, router)
	}
}

func (handler *Handler) GetNoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"res": "404"})
}
