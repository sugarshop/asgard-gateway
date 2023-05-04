package handler

import (
	"github.com/gin-gonic/gin"
)

type OpenAIHandler struct {
}

func NewOpenAIHandler() *OpenAIHandler {
	return &OpenAIHandler{}
}

func (h *OpenAIHandler) Register(e *gin.Engine)  {
	e.POST("/v1/openai/chat/completions", JSONWrapper(h.Completions))
}

func (h *OpenAIHandler) Completions(c *gin.Context) (interface{}, error) {
	//ctx := util.RPCContext(c)
	return nil, nil
}