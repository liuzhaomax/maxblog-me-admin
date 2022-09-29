package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-admin/internal/middleware/interceptor"
	srcHandler "maxblog-me-admin/src/handler"
)

var APISet = wire.NewSet(wire.Struct(new(Handler), "*"), wire.Bind(new(IHandler), new(*Handler)))

type Handler struct {
	HandlerUser *srcHandler.HUser
}

type IHandler interface {
	Register(app *gin.Engine, itcpt *interceptor.Interceptor)
}

func (handler *Handler) Register(app *gin.Engine, itcpt *interceptor.Interceptor) {
	handler.RegisterRouter(app, itcpt)
}
