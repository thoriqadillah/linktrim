package cache

import (
	"context"
	"sync"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, val []byte) error
}

var instance sync.Map
