package redis

import (
	"context"
	"time"
)

var ctx = context.Background()

func Set(key string, value string, expireTime time.Duration) error {
	return NewClient().Set(ctx, key, value, expireTime).Err()
}

func Get(key string) (string, error) {
	return NewClient().Get(ctx, key).Result()
}
