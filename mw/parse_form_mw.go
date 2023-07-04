package mw

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/util"
)

// ParseFormMiddleware parse form, such as device
func ParseFormMiddleware(c *gin.Context) {
	ctx := util.RPCContext(c)
	if err := c.Request.ParseForm(); err != nil {
		log.Println(ctx, "parse form failed ", err)
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(ctx, "read request body error: ", err)
	}
	// rewrite body after read it.
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	c.Next()
}
