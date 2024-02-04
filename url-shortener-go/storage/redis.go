package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var (
	redisOnce sync.Once
	redisCli  *redis.Client
)

// RedisClient returns a singleton instance of the Redis client
func RedisClient() *redis.Client {
	redisOptions := &redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
		DB:   0,
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword != "" {
		redisOptions.Password = redisPassword
	}

	redisOnce.Do(func() {
		redisCli = redis.NewClient(redisOptions)
	})

	return redisCli
}

func RedisSet(key string, value interface{}, ttl time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = RedisClient().Set(context.Background(), key, jsonValue, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisGet(key string) (map[string]interface{}, error) {
	val, err := RedisClient().Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
