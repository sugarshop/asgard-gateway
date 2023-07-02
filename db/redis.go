package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/sugarshop/env"
)

// InitRedisClient 初始化redis client
func InitRedisClient(ctx context.Context) {
	redisConfStr, ok := env.GlobalEnv().Get("REDIS")
	if !ok || len(redisConfStr) == 0 {
		panic("no configuration for redis")
	}
	rc = initRedisClientWithConfStr(redisConfStr)
	// test connection
	err := rc.Get(ctx, "----not_existed_key----").Err()
	if err != redis.Nil {
		panic(err)
	}
	log.Printf("connect to redis success, server: %s", redisConfStr)
}

func initRedisClientWithConfStr(confStr string) *redis.Client {
	// redisConfStr: redis://<user>:<password>@<host>:<port>/<db_number>
	conf, err := redis.ParseURL(confStr)
	if err != nil {
		panic(fmt.Errorf("parse REDIS config failed: %+v", err))
	}
	return redis.NewClient(conf)
}

var rc *redis.Client

// RedisClient returns a redis client
func RedisClient() *redis.Client {
	return rc
}
