package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	res := CompletionDB()
	assert.NotNil(t, res)
}