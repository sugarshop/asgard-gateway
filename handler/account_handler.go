package handler

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/model"
	"github.com/sugarshop/asgard-gateway/service"
	"github.com/sugarshop/asgard-gateway/util"
)

type AccountHandler struct {
}

// NewAccountHandler return payment handler
func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (h *AccountHandler) Register(e *gin.Engine) {
	e.POST("/v1/account/clerk/webhook", JSONWrapper(h.AccountWebHook))
}

// AccountWebHook webhook
func (h *AccountHandler) AccountWebHook(c *gin.Context) (interface{}, error) {
	var reqBody model.ClerkEvent
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
	header := c.Request.Header
	if err := service.ClerkServiceInstance().ListenClerkWebHook(ctx, &reqBody, rawBody, header); err != nil {
		fmt.Printf("[AccountWebHook]: ListenClerkWebHook err:%s\n", err)
		return nil, err
	}
	return map[string]interface{}{}, nil
}
