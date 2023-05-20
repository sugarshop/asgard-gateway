package model

import (
	"github.com/stretchr/testify/assert"
	"github.com/sugarshop/asgard-gateway/db"
	"testing"
)

func TestCompletion_Save(t *testing.T) {
	db.Init()
	compl := Completion{
		ChatID:    "chat-001",
		Model:     "gpt-3.5-turbo-0301",
		Content:   "i love u",
	}
	err := compl.Save()

	assert.NotNil(t, err)
}