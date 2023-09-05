package rdb

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	Context context.Context
	Client  *redis.Client
}
