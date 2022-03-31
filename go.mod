module uims

go 1.16

require (
	github.com/Shopify/sarama v1.30.1
	github.com/envoyproxy/protoc-gen-validate v0.6.3
	github.com/go-kratos/kratos/contrib/metrics/prometheus/v2 v2.0.0-20220330020930-99a0646acb98
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20211119091424-ef3322ec0764
	github.com/go-kratos/kratos/v2 v2.2.1
	github.com/go-redis/redis/extra/redisotel v0.3.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/google/wire v0.5.0
	github.com/hashicorp/consul/api v1.11.0
	github.com/jinzhu/copier v0.3.5
	github.com/prometheus/client_golang v1.9.0
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/objx v0.2.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama v0.28.0
	go.opentelemetry.io/otel v1.6.1
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.6.1
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.6.1
	go.opentelemetry.io/otel/sdk v1.6.1
	go.opentelemetry.io/otel/trace v1.6.1
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa
	golang.org/x/net v0.0.0-20220114011407-0dd24b26b47d // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	google.golang.org/genproto v0.0.0-20220126215142-9970aeb2e350
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.0.6
	gorm.io/gorm v1.21.9
)
