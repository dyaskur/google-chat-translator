package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var ctx = context.Background()

var redisURI = os.Getenv("REDIS_URI")

var addr, _ = redis.ParseURL(redisURI)

var rdb = redis.NewClient(addr)

func SetCache(key string, val string) {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		println(err)
		panic(err)
	}
}

func GetCache(key string) (string, error) {
	value, err := rdb.Get(ctx, key).Result()
	//if errors.Is(err, redis.Nil) {
	//	// not exist
	//	return ""
	//} else if err != nil {
	//	return ""
	//}
	return value, err
}
