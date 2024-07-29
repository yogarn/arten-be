package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/yogarn/arten/pkg/config"
)

func NewRedisClient() *redis.Client {
	host, password := config.LoadRedisCredentials()
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}
