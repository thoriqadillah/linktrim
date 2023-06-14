package db

import (
	"log"

	"github.com/redis/go-redis/v9"
)

var rediskey = "redis"

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
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	store.Store(k, client)
}
