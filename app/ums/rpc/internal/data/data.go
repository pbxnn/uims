package data

import (
	"uims/app/ums/rpc/internal/conf"
	"uims/pkg/kafka"
	gp "uims/third_party/gorm_plugin"

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
	NewUserRepo,
)

// Data .
type Data struct {
	log   *log.Helper
	db    *gorm.DB
	cache *redis.Client
	kp    *kafka.KafkaPubClient
}

func (d Data) GetDB() *gorm.DB {
	return d.db
}

// NewData .
func NewData(db *gorm.DB, cache *redis.Client, kp *kafka.KafkaPubClient, logger log.Logger) (*Data, func(), error) {
	d := &Data{
		log:   log.NewHelper(log.With(logger, "module", "user-service/data")),
		db:    db,
		cache: cache,
		kp:    kp,
	}
	cleanup := func() {}

	return d, cleanup, nil
}

func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(logger)
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
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
