package cache

import (
	"context"
	"sync"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, val []byte, exp time.Duration) error
	Delete(ctx context.Context, key ...string) error
}

var once sync.Once
var instance Cache
var caches map[string]Cache

const REDIS = "redis"

func New(provider ...string) Cache {
	k := REDIS
	if len(provider) > 0 {
		k = provider[0]
	}

	if instance == nil {
		instance = caches[k]
		return caches[k]
	}

	return instance
}

func register(name string, implementation Cache) {
	once.Do(func() {
		caches = make(map[string]Cache)
	})

	caches[name] = implementation
}
