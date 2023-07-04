package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisSet(t *testing.T) {
	ctx := context.Background()
	MockRedis(t)

	err := RedisClient().HSet(ctx, "openai_api_key_test", "123456", "1").Err()
	assert.Equalf(t, nil, err, "")
	assert.Equal(t, true, rc.HExists(ctx, "openai_api_key_test", "123456").Val())
	assert.Equal(t, false, rc.HExists(ctx, "openai_api_key_test", "1234567").Val())
}
