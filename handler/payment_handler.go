package handler

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/model"
	"github.com/sugarshop/asgard-gateway/service"
	"github.com/sugarshop/asgard-gateway/util"
)

type PaymentHandler struct {
}

// NewPaymentHandler return payment handler
func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (h *PaymentHandler) Register(e *gin.Engine) {
	e.GET("/v1/payment/lemonsqueezy/checkout", JSONWrapper(h.LemonSqueezy))
	e.POST("/v1/payment/lemonsqueezy/webhook", JSONWrapper(h.WebHook))
}

// LemonSqueezy create lemonsqueezy checkout link
func (s *PaymentHandler) LemonSqueezy(c *gin.Context) (interface{}, error) {
	// todo: 防止高频攻击，需要进行限频.
	ctx := context.Background()
	link, err := service.LemonSqueezyServiceInstance().CreateCheckoutLink(ctx, "uiduiduid")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"link": link,
	}, nil
}

func (h *PaymentHandler) WebHook(c *gin.Context) (interface{}, error) {
	var reqBody model.LemonSqueezyRequest
	ctx := util.RPCContext(c)
	// bind json to reqBody
	rawBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("[WebHook]: ioutil.ReadAll err: ", err)
		return nil, err
	}
	// rewrite data into body, ioutil.ReadAll will clear data in c.Request.Body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	signature := c.GetHeader("X-Signature")
	if err := service.WebHookServiceInstance().ListenLemonSqueezy(ctx, signature, &reqBody, rawBody); err != nil {
		fmt.Printf("[WebHook]: ListenLemonSqueezy err:%s\n", err)
		return nil, err
	}
	return map[string]interface{}{}, nil
}
