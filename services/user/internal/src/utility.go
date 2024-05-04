package src

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Utility struct {
	Memcache *redis.Client
}

func (u *Utility) Store(ctx context.Context, key string, value any) error {
	stringifiedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = u.Memcache.Set(ctx, key, stringifiedValue, time.Hour*72).Err()
	if err != nil {
		return err
	}

	return nil
}

func (u *Utility) Retrieve(ctx context.Context, key string, result any) error {
	value, err := u.Memcache.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		return err
	}

	return nil
}
