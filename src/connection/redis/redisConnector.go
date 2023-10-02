package redis

import (
	"context"
	"corporateTest/src/helpers"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var log = helpers.GetLogger()

// RedisInstance func
func RedisInstance() *redis.Client {

	// Load environment variables
	redisUrl, err := helpers.GetEnvStringVal("REDIS_URL")
	if err != nil {
		log.Error("Failed to load environment variable : REDIS_URL")
		os.Exit(1)
	}

	redisPassword, err := helpers.GetEnvStringVal("REDIS_PASSWORD")
	if err != nil {
		log.Error("Failed to load environment variable : REDIS_PASSWORD")
		log.Debug(err.Error())
		os.Exit(1)
	}

	log.Info("Connecting to Redis Instance : " + redisUrl)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	return rdb
}

// Client Database instance
var redisClient *redis.Client = RedisInstance()
var ctx = context.Background()

func Set(key string, value string) {
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Error("Error in Redis Set.")
		log.Debug(err.Error())
	}
}

func HSet(key string, values []string) {
	err := redisClient.HSet(ctx, key, values).Err()
	if err != nil {
		log.Error("Error in Redis HSet.")
		log.Debug(err.Error())
	}
}

func HGetAll(key string) *redis.StringStringMapCmd {
	return redisClient.HGetAll(ctx, key)

}

func Get(key string) (string, error) {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func SetWithExp(key string, value string, exp time.Duration) error {
	err := redisClient.Set(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func Delete(key string) (int64, error) {
	val, err := redisClient.Del(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func GetAllKeys() ([]string, error) {
	keys, err := redisClient.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}
