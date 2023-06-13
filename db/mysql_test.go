package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sugarshop/env"
)

func TestMysqlInit(t *testing.T) {
	env.LoadGlobalEnv("../conf/test.json")
	MysqlInit()
	db := SugarShopDB()
	assert.NotNil(t, db)
}
