package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheMiddleware struct {
	client *redis.Client
}

func NewCacheMiddleware(client *redis.Client) *CacheMiddleware {
	return &CacheMiddleware{client: client}
}

func (cm *CacheMiddleware) GetFromCache(key string, dest interface{}) error {
	val, err := cm.client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func (cm *CacheMiddleware) SetToCache(key string, value interface{}, expiration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return cm.client.Set(context.Background(), key, jsonValue, expiration).Err()
}

func (cm *CacheMiddleware) DeleteFromCache(key string) error {
	return cm.client.Del(context.Background(), key).Err()
}

func (cm *CacheMiddleware) ClearUserCache(userID uint) error {
	keys := []string{
		fmt.Sprintf("user:%d", userID),
		"users:all",
	}
	return cm.client.Del(context.Background(), keys...).Err()
}
