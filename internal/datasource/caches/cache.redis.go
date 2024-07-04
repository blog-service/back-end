package caches

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
	Del(key string) error
}

type redisCache struct {
	host     string
	db       int
	password string
	expires  time.Duration
	client   *redis.Client
}

func NewRedisCache(host string, db int, password string, expires time.Duration) RedisCache {
	return &redisCache{
		host:     host,
		db:       db,
		password: password,
		expires:  expires,
		client: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       db,
		}),
	}
}

func (cache *redisCache) Set(key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return cache.client.Set(ctx, key, jsonData, cache.expires*time.Minute).Err()
}

func (cache *redisCache) Get(key string) (email string, err error) {
	ctx := context.Background()
	val, err := cache.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(val), &email)
	return email, err
}

func (cache *redisCache) Del(key string) error {
	ctx := context.Background()
	return cache.client.Del(ctx, key).Err()
}
