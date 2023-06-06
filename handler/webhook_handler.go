package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/model"
	"net/http"
)

type WebhookHandler struct {
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

func (h *WebhookHandler) Register(e *gin.Engine)  {
	e.POST("/webhook/lemonsqueezy", JSONWrapper(h.Lemonsqueezy))
}

func (h *WebhookHandler) Lemonsqueezy(c *gin.Context) (interface{}, error) {
	// todo fixme
	var reqBody model.CompletionsReqBody
	// bind json to reqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	return nil, nil
}