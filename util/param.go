package util

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/model"
)

// String 获取string格式参数
func String(c *gin.Context, key string) (string, error) {
	value := c.Request.Form.Get(key)
	if len(value) == 0 {
		return "", fmt.Errorf("param %s not set, code:%d", key, model.ErrCodeInvalidParam)
	}
	return value, nil
}

// StringWithDefault 获取string格式参数，若未赋值，则返回默认值
func StringWithDefault(c *gin.Context, key, defaultValue string) string {
	value, err := String(c, key)
	if err != nil {
		return defaultValue
	}
	return value
}

// Int 获取int格式参数
func Int(c *gin.Context, key string) (int, error) {
	value := c.Request.Form.Get(key)
	if len(value) == 0 {
		return 0, fmt.Errorf("param %s not set, code: %d", key, model.ErrCodeInvalidParam)
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}
	return v, nil
}

// PositiveInt 获取大于0的int格式参数
func PositiveInt(c *gin.Context, key string) (int, error) {
	v, err := Int(c, key)
	if err != nil {
		return 0, err
	}
	if v <= 0 {
		return 0, fmt.Errorf("param %s value %d invalid, code: %d", key, v, model.ErrCodeInvalidParam)
	}
	return v, nil
}

// NonNegativeInt 获取不小于0的int格式参数
func NonNegativeInt(c *gin.Context, key string) (int, error) {
	v, err := Int(c, key)
	if err != nil {
		return 0, err
	}
	if v < 0 {
		return 0, fmt.Errorf("param %s value %d invalid, code: %d", key, v, model.ErrCodeInvalidParam)
	}
	return v, nil
}

// Bool 获取bool格式参数
func Bool(c *gin.Context, key string) (bool, error) {
	value := c.Request.Form.Get(key)
	if len(value) == 0 {
		return false, fmt.Errorf("param %s not set, code: %d", key, model.ErrCodeInvalidParam)
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("err: %w, code: %d", err, model.ErrCodeInvalidParam)
	}
	return v, nil
}

// BoolWithDefault 获取bool格式参数，未赋值则返回默认
func BoolWithDefault(c *gin.Context, key string, defaultValue bool) bool {
	value, err := Bool(c, key)
	if err != nil {
		return defaultValue
	}
	return value
}

// Int64 获取int64格式参数
func Int64(c *gin.Context, key string) (int64, error) {
	value := c.Request.Form.Get(key)
	if len(value) == 0 {
		return 0, fmt.Errorf("param %s not set, code: %d", key, model.ErrCodeInvalidParam)
	}
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error: %w, code: %d", err, model.ErrCodeInvalidParam)
	}
	return v, nil
}

// Int64WithDefault 获取int64格式参数，未赋值返回默认值
func Int64WithDefault(c *gin.Context, key string, defaultValue int64) int64 {
	value, err := Int64(c, key)
	if err != nil {
		return defaultValue
	}
	return value
}

// Int32 获取int32格式参数
func Int32(c *gin.Context, key string) (int32, error) {
	value := c.Request.Form.Get(key)
	if len(value) == 0 {
		return 0, fmt.Errorf("param %s not set, code: %d", key, model.ErrCodeInvalidParam)
	}
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error: %w, code: %d", err, model.ErrCodeInvalidParam)
	}
	return int32(v), nil
}

// Int32WithDefault 获取int32格式参数，若未赋值则返回默认值
func Int32WithDefault(c *gin.Context, key string, defaultValue int32) int32 {
	value, err := Int32(c, key)
	if err != nil {
		return defaultValue
	}
	return value
}

// PositiveInt64 获取大于0的int64格式参数
func PositiveInt64(c *gin.Context, key string) (int64, error) {
	v, err := Int64(c, key)
	if err != nil {
		return 0, err
	}
	if v <= 0 {
		return 0, fmt.Errorf("param %s value %d invalid, code: %d", key, v, model.ErrCodeInvalidParam)
	}
	return v, nil
}

//func isApp(c *gin.Context) bool {
//	return strings.Contains(c.GetHeader("User-Agent"), "is_app")
//}
//
//func isMiniProgram(c *gin.Context) bool {
//	// 如果包含 micro messenger，则代表是微信内，包含 header 则代表是小程序的请求
//	if strings.Contains(c.GetHeader("User-Agent"), "MicroMessenger") && strings.Contains(c.GetHeader("Is-WebApp"), "weixin-mini-program") {
//		return true
//	}
//	return false
//}
