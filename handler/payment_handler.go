package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/service"
)

type PaymentHandler struct {
}

// NewPaymentHandler return payment handler
func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (h *PaymentHandler) Register(e *gin.Engine) {
	e.GET("/v1/payment/lemonsqueezy/checkout", JSONWrapper(h.LemonSqueezy))
}

// LemonSqueezy create lemonsqueezy checkout link
func (s *PaymentHandler) LemonSqueezy(c *gin.Context) (interface{}, error) {
	ctx := context.Background()
	link, err := service.LemonSqueezyServiceInstance().CreateCheckoutLink(ctx, "uiduiduid")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"link": link,
	}, nil
}
