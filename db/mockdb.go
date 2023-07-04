package db

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sugarshop/sugarredis"
)

// MockRedis test redis server in memory for ut.
func MockRedis(t *testing.T) {
	mr := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	rc = &sugarredis.SgRedisClient{}
	rc.SetClients(client)
}
