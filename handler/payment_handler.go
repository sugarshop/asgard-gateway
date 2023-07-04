package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/service"
	"github.com/sugarshop/asgard-gateway/util"
	lemonsqueezy "github.com/sugarshop/lemonsqueezy-go"
)

type PaymentHandler struct {
}

// NewPaymentHandler return payment handler
func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (h *PaymentHandler) Register(e *gin.Engine) {
	e.GET("/v1/payment/lemonsqueezy/checkout", JSONWrapper(h.CreateCheckout))
	e.POST("/v1/payment/lemonsqueezy/webhook", JSONWrapper(h.WebHook))
}

// CreateCheckout create lemonsqueezy checkout link
func (s *PaymentHandler) CreateCheckout(c *gin.Context) (interface{}, error) {
	// todo: 防止高频攻击，需要进行限频.
	ctx := context.Background()
	uid, err := util.String(c, "uid")
	if err != nil {
		log.Println("[CreateCheckout]: parser uid err ", err)
		return nil, err
	}
	link, err := service.LemonSqueezyServiceInstance().CreateCheckoutLink(ctx, uid)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"link": link,
	}, nil
}

func (h *PaymentHandler) WebHook(c *gin.Context) (interface{}, error) {
	var reqBody lemonsqueezy.WebhookRequest
	ctx := util.RPCContext(c)
	// bind json to reqBody
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("[AccountWebHook]: ioutil.ReadAll err: ", err)
		return nil, err
	}
	// rewrite data into body, ioutil.ReadAll will clear data in c.Request.Body
	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	signature := c.GetHeader("X-Signature")
	if err := service.LemonSqueezyServiceInstance().ListenWebhook(ctx, signature, &reqBody, rawBody); err != nil {
		fmt.Printf("[AccountWebHook]: ListenWebhook err:%s\n", err)
		return nil, err
	}
	return map[string]interface{}{}, nil
}
