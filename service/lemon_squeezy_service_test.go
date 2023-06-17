package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sugarshop/env"
)

func TestLemonSqueezyService_ListCustomers(t *testing.T) {
	env.LoadGlobalEnv("../conf/test.json")
	ctx := context.Background()
	apiResponse, err := LemonSqueezyServiceInstance().ListCustomers(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, apiResponse)
}

func TestLemonSqueezyService_CreateCheckout(t *testing.T) {
	env.LoadGlobalEnv("../conf/test.json")
	ctx := context.Background()
	link, err := LemonSqueezyServiceInstance().CreateCheckoutLink(ctx, "unit_test")
	assert.Nil(t, err)
	assert.NotNil(t, link)
}
