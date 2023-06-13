package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemonSqueezyService_ListCustomers(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := LemonSqueezyServiceInstance().ListCustomers(ctx)
	fmt.Println(apiResponse, err)
	assert.Nil(t, err)
}

func TestLemonSqueezyService_CreateCheckout(t *testing.T) {
	ctx := context.Background()
	link, err := LemonSqueezyServiceInstance().CreateCheckoutLink(ctx, "ssddd")
	fmt.Println(link, err)
	assert.Nil(t, err)
}
