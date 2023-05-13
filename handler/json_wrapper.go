package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/kitex_gen/base"
)

// JSONWrapper Encapsulate the data processing function as a JSON API return; pay attention to using PureJSON to avoid Gin performing HTML escaping during serialization of data.
func JSONWrapper(fn func(*gin.Context) (interface{}, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		data, err := fn(c)

		// if write is been writen at here, do not write again or you will get panic
		if c.Writer.Written() {
			return
		}

		if err != nil {
			//c.Set(tracing.CtxRespCodeKey, base.FAILED)
			c.PureJSON(http.StatusOK, &ErrResp{
				Code:   base.FAILED,
				Msg:    err.Error(),
				Detail: err.Error(),
			})
			return
		}

		c.PureJSON(http.StatusOK, &DataResp{
			Code: base.OK,
			Data: data,
		})
	}
}

// StreamWrapper Encapsulate the data processing function as a STREAM API return; pay attention to using PureJSON to avoid Gin performing HTML escaping during serialization of data.
func StreamWrapper(fn func(*gin.Context) error) func(*gin.Context) {
	return func(c *gin.Context) {
		err := fn(c)

		// if write is been writen at here, do not write again or you will get panic
		if c.Writer.Written() {
			return
		}

		if err != nil {
			//c.Set(tracing.CtxRespCodeKey, base.FAILED)
			c.PureJSON(http.StatusOK, &ErrResp{
				Code:   base.FAILED,
				Msg:    err.Error(),
				Detail: err.Error(),
			})
			return
		}
	}
}

type ErrResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
}

type DataResp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
