package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	umsService "uims/api/ums/rpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserServiceClient, NewUserRepo)

// Data .
type Data struct {
	usc    umsService.UserClient
	logger log.Logger
}

// NewData .
func NewData(usc umsService.UserClient, logger log.Logger) (*Data, func(), error) {
	d := &Data{usc: usc, logger: logger}
	return d, nil, nil
}

func NewUserServiceClient(r registry.Discovery) umsService.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///uims.ums.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := umsService.NewUserClient(conn)
	return c
}
