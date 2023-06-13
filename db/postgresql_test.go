package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sugarshop/env"
)

func TestPostgreSqlInit(t *testing.T) {
	env.LoadGlobalEnv("../conf/test.json")
	PostgreSqlInit()
	res := CompletionDB()
	assert.NotNil(t, res)
}
