package data

import (
	"gorm.io/gorm/schema"
	"uims/app/orgms/rpc/internal/conf"
	gp "uims/third_party/gorm_plugin"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewCache,
	NewKafkaAsyncProducer,
	NewKafkaSyncProducer,
	NewCompanyRepo,
)

// Data .
type Data struct {
	db    *gorm.DB
	cache *redis.Client
	kap   sarama.AsyncProducer
	ksp   sarama.SyncProducer
}

// NewData .
func NewData(db *gorm.DB, cache *redis.Client, kap sarama.AsyncProducer, ksp sarama.SyncProducer, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		kap.Close()
		ksp.Close()
	}
	return &Data{db: db, cache: cache, kap: kap, ksp: ksp}, cleanup, nil
}

func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(logger)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
		return db
	}

	log.Infof("init DB...")
	err = db.Use(&gp.OpentracingPlugin{})
	if err != nil {
		log.Fatal(err)
		return db
	}

	log.Infof("init DB success...")
	return db
}

func NewCache(c *conf.Data, logger log.Logger) *redis.Client {
	cache := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	cache.AddHook(redisotel.TracingHook{})

	return cache
}

func NewKafkaAsyncProducer(conf *conf.Data) sarama.AsyncProducer {
	c := sarama.NewConfig()
	c.Producer.Return.Successes = false
	c.Producer.Return.Errors = true
	p, err := sarama.NewAsyncProducer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}

	return p
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
