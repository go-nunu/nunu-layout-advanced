package repository

import (
	"context"
	"fmt"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"time"
)

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(db *gorm.DB, rdb *redis.Client, logger *log.Logger) *Repository {
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}

type ctxTransactionKey struct {
}

func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTransactionKey{})
	if v == nil {
		return r.db.WithContext(ctx)
	} else {
		tx, ok := v.(*gorm.DB)
		if !ok {
			panic(fmt.Sprintf("invalid transaction type: %T", v))
		}
		return tx
	}
}

func (r *Repository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	ctx = context.WithValue(ctx, ctxTransactionKey{}, tx)
	if err := fn(ctx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func NewDB(conf *viper.Viper, l *log.Logger) *gorm.DB {
	logger := zapgorm2.New(l.Logger)
	logger.SetAsDefault()
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{Logger: logger})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	return db
}
func NewRedis(conf *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("data.redis.addr"),
		Password: conf.GetString("data.redis.password"),
		DB:       conf.GetInt("data.redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}
