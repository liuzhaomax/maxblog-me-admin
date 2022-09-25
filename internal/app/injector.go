package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-admin/internal/api"
	"maxblog-me-admin/internal/middleware/interceptor"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine      *gin.Engine
	Handler     *api.Handler
	Interceptor *interceptor.Interceptor
}
