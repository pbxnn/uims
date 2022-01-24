package main

import (
	"uims/app/orgms/rpc/internal/biz"
	"uims/app/orgms/rpc/internal/conf"
	"uims/app/orgms/rpc/internal/data"
	"uims/app/orgms/rpc/internal/server"
	"uims/app/orgms/rpc/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}
