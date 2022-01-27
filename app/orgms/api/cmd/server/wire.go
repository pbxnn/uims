// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"uims/app/orgms/api/internal/biz"
	"uims/app/orgms/api/internal/conf"
	"uims/app/orgms/api/internal/data"
	"uims/app/orgms/api/internal/server"
	"uims/app/orgms/api/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.KafkaProducer, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}
