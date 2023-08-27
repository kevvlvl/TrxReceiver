package rdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var client *redis.Client
var ctx = context.Background()

func Client() *redis.Client {

	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	}

	return client
}

func Get(key string, client *redis.Client) string {
	v, err := client.Get(ctx, key).Result()

	if err != nil {
		log.Error().Msgf("Error trying to get %s, %v", key, err)
	}

	return v
}

func Set(key string, value string, client *redis.Client) {
	err := client.Set(ctx, key, value, 0).Err()

	if err != nil {
		log.Error().Msgf("Error trying to set %s, %v", key, err)
	}
}
