package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/thoriqadillah/linktrim/lib/env"
)

// TODO: get to know how to update the value without updating the TTL
type redisCache struct {
	client *redis.Client
}

var (
	addr     = env.Get("REDIS_ADDR").ToString()
	password = env.Get("REDIS_PASSWORD").ToString()
	dbNum    = env.Get("REDIS_DB").ToInt()
)

func Redis() Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbNum,
	})

	return &redisCache{
		client: client,
	}
}

func (c *redisCache) Get(ctx context.Context, key string) ([]byte, error) {
	res, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("user not found: %s", err.Error())
	}

	return res, nil
}

func (c *redisCache) Set(ctx context.Context, key string, value []byte, duration ...time.Duration) error {
	exp := time.Duration(0)
	if len(duration) > 0 {
		exp = duration[0]
	}

	if err := c.client.Set(ctx, key, value, exp).Err(); err != nil {
		return fmt.Errorf("could not store into cache: %s", err.Error())
	}

	return nil
}

func (c *redisCache) Delete(ctx context.Context, key ...string) error {
	return c.client.Del(ctx, key...).Err()
}

func init() {
	register(REDIS, Redis())
}
