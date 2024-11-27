package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var ctx = context.Background()

var rdb *redis.Client

func client() *redis.Client {
	if rdb != nil {
		return rdb
	}
	redisURI, ok := os.LookupEnv("REDIS_URI")
	if !ok {
		println("REDIS_URI is not set")
		redisURI = "redis://localhost:6379" // fallback to default URI if not set
	}
	var addr, _ = redis.ParseURL(redisURI)
	rdb = redis.NewClient(addr)
	return rdb
}

func SetCache(key string, val string) {
	err := client().Set(ctx, key, val, 0).Err()
	if err != nil {
		println(err)
		panic(err)
	}
}

func GetCache(key string) (string, error) {
	value, err := client().Get(ctx, key).Result()
	//if errors.Is(err, redis.Nil) {
	//	// not exist
	//	return ""
	//} else if err != nil {
	//	return ""
	//}
	return value, err
}
