package remote

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sugarshop/env"
)

func TestLemonSqueezyService_ListVariants(t *testing.T) {
	env.LoadGlobalEnv("../conf/test.json")
	ctx := context.Background()
	response, err := LemonSqueezyServiceInstance().ListVariants(ctx)
	assert.NotNil(t, response)
	assert.Nil(t, err)
}
