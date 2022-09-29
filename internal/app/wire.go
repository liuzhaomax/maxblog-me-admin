//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"maxblog-me-admin/internal/api"
	"maxblog-me-admin/internal/conf"
	"maxblog-me-admin/internal/core"
	"maxblog-me-admin/internal/middleware/interceptor"
	"maxblog-me-admin/src/handler"
	"maxblog-me-admin/src/service"
)

func InitInjector() (*Injector, error) {
	wire.Build(
		conf.InitGinEngine,
		api.APISet,
		interceptor.InterceptorSet,
		interceptor.AuthSet,
		core.ResponseSet,
		core.LoggerSet,
		handler.HandlerSet,
		service.ServiceSet,
		InjectorSet,
	)
	return new(Injector), nil
}
