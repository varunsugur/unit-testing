package database

import (
	"golang/config"

	"github.com/redis/go-redis/v9"
)

func ConnectToRedis(cfg config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.RedisPost,
		Password: cfg.RedisConfig.RedisPassword, //no password set
		DB:       cfg.RedisConfig.RedisDB,       //use default db
	})
	return rdb
}
