package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
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

// IncrementCounter adds n to the counter at key and keeps it for ttl.
func IncrementCounter(key string, n int64, ttl time.Duration) (int64, error) {
	pipe := client().TxPipeline()
	incr := pipe.IncrBy(ctx, key, n)
	pipe.Expire(ctx, key, ttl)
	if _, err := pipe.Exec(ctx); err != nil {
		return 0, err
	}
	return incr.Val(), nil
}
