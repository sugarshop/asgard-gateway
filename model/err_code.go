package model

import "fmt"

const (
	//  general error code

	// ErrCodeInvalidParam param code error
	ErrCodeInvalidParam = 1

	// ErrCodeDataNotFound data not found
	ErrCodeDataNotFound = 2

	// ErrCodeInternalServerError system internal error
	ErrCodeInternalServerError = 500
)

// ErrDataNotFound 无有效数据
var ErrDataNotFound = fmt.Errorf("no valid data, code: %d", ErrCodeDataNotFound)
