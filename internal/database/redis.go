package database

import "github.com/redis/go-redis/v9"

func ConnectToRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //no password set
		DB:       0,  //use default db
	})
	return rdb
}
