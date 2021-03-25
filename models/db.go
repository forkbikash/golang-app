package models

import "github.com/go-redis/redis"

var client *redis.Client

func Init() {
	// creating a redis client
	client = redis.NewClient(&redis.Options{
		Addr : "localhost:6379", // 6379 is the default port for redis
	})
}
