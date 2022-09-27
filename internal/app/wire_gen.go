// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"maxblog-me-admin/internal/api"
	"maxblog-me-admin/internal/conf"
	"maxblog-me-admin/internal/core"
	"maxblog-me-admin/internal/middleware/interceptor"
	"maxblog-me-admin/src/handler"
	"maxblog-me-admin/src/service"
)

// Injectors from wire.go:

func InitInjector() (*Injector, error) {
	bUser := &service.BUser{}
	logger := &core.Logger{}
	response := &core.Response{
		ILogger: logger,
	}
	hUser := &handler.HUser{
		BUser: bUser,
		IRes:  response,
	}
	apiHandler := &api.Handler{
		HandlerUser: hUser,
	}
	engine := conf.InitGinEngine(apiHandler)
	interceptorInterceptor := &interceptor.Interceptor{}
	injector := &Injector{
		Engine:      engine,
		Handler:     apiHandler,
		Interceptor: interceptorInterceptor,
	}
	return injector, nil
}
