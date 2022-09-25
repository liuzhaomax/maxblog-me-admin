//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"maxblog-me-template/internal/api"
	"maxblog-me-template/internal/conf"
	"maxblog-me-template/internal/core"
	"maxblog-me-template/internal/middleware/interceptor"
	dataHandler "maxblog-me-template/src/handler"
	dataService "maxblog-me-template/src/service"
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
