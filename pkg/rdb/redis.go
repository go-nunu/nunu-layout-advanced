package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

func NewRedis(conf *viper.Viper) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("data.redis.addr"),
		Password: conf.GetString("data.redis.password"),
		DB:       conf.GetInt("data.redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}
