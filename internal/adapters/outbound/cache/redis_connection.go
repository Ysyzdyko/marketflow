package cache

import (
	"context"
	"fmt"
	"marketflow/internal/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

type redisCache struct {
	client *redis.Client
}

func NewRedis(cfg *config.RedisConfig) *redisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(fmt.Sprintf("не удалось подключиться к Redis: %v", err))
	}

	return &redisCache{client: client}
}

func (r *redisCache) Set(key string, value string) error {
	return r.client.Set(ctx, key, value, 5*time.Minute).Err()
}

func (r *redisCache) Get(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisCache) MGet(keys ...string) ([]interface{}, error) {
	return r.client.MGet(ctx, keys...).Result()
}

func (r *redisCache) Keys(pattern string) ([]string, error) {
	return r.client.Keys(ctx, pattern).Result()
}

func (r *redisCache) FlushAll() error {
	return r.client.FlushAll(ctx).Err()
}
