package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/model"
	"github.com/sugarshop/asgard-gateway/service"
	"github.com/sugarshop/asgard-gateway/util"
)

type WebhookHandler struct {
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

func (h *WebhookHandler) Register(e *gin.Engine) {
	e.POST("/webhook/lemonsqueezy", JSONWrapper(h.LemonSqueezy))
}

func (h *WebhookHandler) LemonSqueezy(c *gin.Context) (interface{}, error) {

	// todo verify X-Signature in header to assure request is from LemonSqueezy

	var reqBody model.LemonSqueezyRequest
	ctx := util.RPCContext(c)
	// bind json to reqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	if err := service.WebHookServiceInstance().ListenLemonSqueezy(ctx, &reqBody); err != nil {
		fmt.Printf("[LemonSqueezy]: ListenLemonSqueezy err:%s\n", err)
		return nil, err
	}
	return map[string]interface{}{}, nil
}
