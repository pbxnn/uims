// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"uims/app/ums/api/internal/biz"
	"uims/app/ums/api/internal/conf"
	"uims/app/ums/api/internal/data"
	"uims/app/ums/api/internal/server"
	"uims/app/ums/api/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	//tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
