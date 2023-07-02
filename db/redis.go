package db

import (
	"context"

	"github.com/sugarshop/env"
	"github.com/sugarshop/sugarredis"
)

// InitRedisClient 初始化redis client
func InitRedisClient(ctx context.Context) {
	redisConfStr, ok := env.GlobalEnv().Get("REDIS")
	if !ok || len(redisConfStr) == 0 {
		panic("no configuration for redis")
	}
	rc = sugarredis.InitRedisClient(ctx, redisConfStr)
}

var rc *sugarredis.SgRedisClient

// RedisClient returns a redis client
func RedisClient() *sugarredis.SgRedisClient {
	return rc
}
