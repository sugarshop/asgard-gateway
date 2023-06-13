package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sugarshop/asgard-gateway/db"
	"github.com/sugarshop/env"
)

func TestCompletion_Save(t *testing.T) {
	env.LoadGlobalEnv("../conf/test.json")
	db.Init()
	compl := Completion{
		ChatID:  "chat-001",
		Model:   "gpt-3.5-turbo-0301",
		Content: "i love u",
		Role:    "test",
	}
	err := compl.Save()

	assert.Nil(t, err)
}
