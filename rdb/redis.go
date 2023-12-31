package rdb

import (
	"TrxReceiver/env"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"time"
)

func Instance() RedisDB {

	return GetRedisClient(env.RedisHost(), env.RedisPort(), env.RedisPass())
}

func GetRedisClient(addr string, port int64, password string) RedisDB {
	redisClient := RedisDB{
		Context: context.Background(),
	}

	redisClient.init(addr, port, password)

	return redisClient
}

func (rdb *RedisDB) Get(key string) string {
	v, err := rdb.Client.Get(rdb.Context, key).Result()

	if err != nil {
		log.Error().Msgf("Error trying to get %s, %v", key, err)
	}

	return v
}

func (rdb *RedisDB) Set(key string, value string, exp time.Duration) string {
	status, err := rdb.Client.Set(rdb.Context, key, value, exp).Result()

	if err != nil {
		log.Error().Msgf("Error trying to set %s, %v", key, err)
		return err.Error()
	}

	log.Debug().Msgf("Set Key successfully: %v", status)
	return status
}

func (rdb *RedisDB) init(addr string, port int64, password string) {
	if rdb.Client == nil {
		rdb.Client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", addr, port),
			Password: password,
			DB:       0,
		})
	}
}
