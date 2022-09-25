package conf

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"maxblog-me-template/internal/api"
)

func init() {
	gin.DefaultWriter = colorable.NewColorableStdout()
	gin.ForceConsoleColor()
}

func InitGinEngine(iHandler api.IHandler) *gin.Engine {
	gin.SetMode(cfg.RunMode) // debug, test, release
	app := gin.Default()
	app.Use(LoggerToFile())
	iHandler.Register(app)
	return app
}
