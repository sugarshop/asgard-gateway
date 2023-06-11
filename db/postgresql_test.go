package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgreSqlInit(t *testing.T) {
	PostgreSqlInit()
	res := CompletionDB()
	assert.NotNil(t, res)
}
