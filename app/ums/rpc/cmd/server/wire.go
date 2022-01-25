// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"uims/app/ums/rpc/internal/biz"
	"uims/app/ums/rpc/internal/conf"
	"uims/app/ums/rpc/internal/data"
	"uims/app/ums/rpc/internal/data/dao"
	"uims/app/ums/rpc/internal/server"
	"uims/app/ums/rpc/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	//tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		dao.ProviderSet,
		newApp,
	))
}
