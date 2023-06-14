package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/thoriqadillah/linktrim/db"
	"github.com/thoriqadillah/linktrim/modules/account/model"
)

var cachekey = "cache:user"

type userCache struct {
	cache *redis.Client
	exp   time.Duration
}

func NewUserCache() Cache {
	val, ok := instance.Load(cachekey)
	if !ok {
		c := &userCache{
			cache: db.Redis(),
			exp:   time.Duration(time.Hour * 24),
		}

		instance.Store(cachekey, c)
		return c
	}

	return val.(*userCache)
}

func (c *userCache) Get(ctx context.Context, key string) (interface{}, error) {
	res, err := db.Redis().Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("user not found: %s", err.Error())
	}

	var val *model.User
	if err := json.Unmarshal(res, &val); err != nil {
		return nil, err
	}

	return val, nil
}

func (c *userCache) Set(ctx context.Context, key string, value []byte) error {
	if err := db.Redis().Set(ctx, key, value, c.exp).Err(); err != nil {
		return fmt.Errorf("could not store into cache: %s", err.Error())
	}

	return nil
}
