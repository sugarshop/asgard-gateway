package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMysqlInit(t *testing.T) {
	MysqlInit()
	db := SugarShopDB()
	assert.NotNil(t, db)
}
