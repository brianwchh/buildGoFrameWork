package plugins

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
	ctx    context.Context
}

func initCache() *Cache {

	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err.Error())
	}

	return &Cache{
		client: client,
		ctx:    ctx,
	}
}

func (cache *Cache) Set(key string, value string, expiration time.Duration) error {
	err := cache.client.Set(cache.ctx, key, value, expiration).Err()

	return err
}

func (cache *Cache) Get(key string) (string, error) {
	val, err := cache.client.Get(cache.ctx, key).Result()

	return val, err
}

func (cache *Cache) Del(key string) error {
	err := cache.client.Del(cache.ctx, key).Err()

	return err
}
