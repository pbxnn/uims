package data

import (
	"context"
	"github.com/Shopify/sarama"

	"uims/api/orgms/rpc"
	"uims/app/orgms/api/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewCompanyClient,
	NewDepartmentClient,
	NewUserClient,
	NewKafkaSyncProducer,
)

// Data .
type Data struct {
	logger           log.Logger
	companyClient    rpc.CompanyClient
	departmentClient rpc.DepartmentClient
	userClient       rpc.UserClient
	syncProducer     sarama.SyncProducer
}

// NewData
func NewData(
	companyClient rpc.CompanyClient,
	departmentClient rpc.DepartmentClient,
	userClient rpc.UserClient,
	syncProducer sarama.SyncProducer,
) (*Data, func(), error) {
	d := &Data{
		companyClient:    companyClient,
		departmentClient: departmentClient,
		userClient:       userClient,
		syncProducer:     syncProducer,
	}
	return d, nil, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewCompanyClient(r registry.Discovery) rpc.CompanyClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///uims.orgms.rpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := rpc.NewCompanyClient(conn)
	return c
}

func NewDepartmentClient(r registry.Discovery) rpc.DepartmentClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///uims.orgms.rpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := rpc.NewDepartmentClient(conn)
	return c
}

func NewUserClient(r registry.Discovery) rpc.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///uims.orgms.rpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := rpc.NewUserClient(conn)
	return c
}

func NewKafkaSyncProducer(conf *conf.Data) sarama.SyncProducer {
	c := sarama.NewConfig()
	c.Producer.Return.Successes = true
	c.Producer.Return.Errors = true

	p, err := sarama.NewSyncProducer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}

	return p
}
