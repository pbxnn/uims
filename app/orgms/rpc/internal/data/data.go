package data

import (
	"fmt"
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
	NewKafkaProducer,
	NewCompanyRepo,
)

// Data .
type Data struct {
	db    *gorm.DB
	cache *redis.Client
	kp    sarama.AsyncProducer
}

// NewData .
func NewData(db *gorm.DB, cache *redis.Client, kp sarama.AsyncProducer, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		kp.Close()
	}
	return &Data{db: db, cache: cache, kp: kp}, cleanup, nil
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

func NewKafkaProducer(conf *conf.Data) sarama.AsyncProducer {
	c := sarama.NewConfig()
	fmt.Println(conf.Kafka.Addrs)
	p, err := sarama.NewAsyncProducer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}

	return p
}
