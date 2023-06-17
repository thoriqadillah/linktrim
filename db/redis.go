package db

import (
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/thoriqadillah/linktrim/lib/env"
)

var rediskey = "redis"
var (
	addr     = env.Get("REDIS_ADDR").ToString()
	password = env.Get("REDIS_PASSWORD").ToString()
	db       = env.Get("REDIS_DB").ToInt()
)

func Redis(key ...string) *redis.Client {
	k := rediskey
	if len(key) > 0 {
		k = key[0]
	}

	client, ok := store.Load(k)
	if !ok {
		log.Panicf("Error: redis client is not initialized")
		return nil
	}

	return client.(*redis.Client)
}

func SetupRedis(key ...string) {
	k := rediskey
	if len(key) > 0 {
		k = key[0]
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	store.Store(k, client)
}
