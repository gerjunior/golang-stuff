package db

import (
	"github.com/go-redis/redis/v9"
)

var Rdb *redis.Client

func RedisInit() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
