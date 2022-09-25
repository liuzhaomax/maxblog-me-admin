//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"maxblog-me-admin/internal/api"
	"maxblog-me-admin/internal/conf"
	"maxblog-me-admin/internal/core"
	"maxblog-me-admin/internal/middleware/interceptor"
	dataHandler "maxblog-me-admin/src/handler"
	dataService "maxblog-me-admin/src/service"
)

func InitInjector() (*Injector, error) {
	wire.Build(
		conf.InitGinEngine,
		api.APISet,
		interceptor.InterceptorSet,
		core.ResponseSet,
		core.LoggerSet,
		dataHandler.HandlerSet,
		dataService.ServiceSet,
		InjectorSet,
	)
	return new(Injector), nil
}
