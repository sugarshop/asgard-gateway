package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemonSqueezyService_ListCustomers(t *testing.T) {
	ctx := context.Background()
	apiResponse, response, err := LemonSqueezyServiceInstance().ListCustomers(ctx)
	fmt.Println(apiResponse, response, err)
	assert.Nil(t, err)
}
