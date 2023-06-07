package service

import (
	"context"
	"sync"

	lemonsqueezy "github.com/NdoleStudio/lemonsqueezy-go"
	"github.com/sugarshop/asgard-gateway/model"
)

type LemonSqueezyService struct {
	Client *lemonsqueezy.Client
}

var (
	lemonSqueezyService *LemonSqueezyService
	lemonSqueezyOnce    sync.Once
)

func LemonSqueezyServiceInstance() *LemonSqueezyService {
	lemonSqueezyOnce.Do(func() {
		lemonSqueezyService = &LemonSqueezyService{
			Client: lemonsqueezy.New(lemonsqueezy.WithAPIKey(model.LemonSqueezyAPIKey)),
		}
	})
	return lemonSqueezyService
}

func (s *LemonSqueezyService) ListCustomers(ctx context.Context) (*lemonsqueezy.CustomersApiResponse, *lemonsqueezy.Response, error) {
	customerApiResponse, response, err := s.Client.Customers.List(ctx)
	return customerApiResponse, response, err
}
