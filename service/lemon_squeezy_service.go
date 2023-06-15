package service

import (
	"context"
	"sync"

	"github.com/sugarshop/asgard-gateway/remote"
	lemonsqueezy "github.com/sugarshop/lemonsqueezy-go"
)

type LemonSqueezyService struct {
}

var (
	lemonSqueezyService *LemonSqueezyService
	lemonSqueezyOnce    sync.Once
)

func LemonSqueezyServiceInstance() *LemonSqueezyService {
	lemonSqueezyOnce.Do(func() {
		lemonSqueezyService = &LemonSqueezyService{}
	})
	return lemonSqueezyService
}

func (s *LemonSqueezyService) ListCustomers(ctx context.Context) (*lemonsqueezy.CustomersApiResponse, error) {
	customerApiResponse, err := remote.LemonSqueezyServiceInstance().ListCustomers(ctx)
	return customerApiResponse, err
}

func (s *LemonSqueezyService) CreateCheckoutLink(ctx context.Context, uid string) (string, error) {
	checkoutApiResponse, err := remote.LemonSqueezyServiceInstance().CreateCheckout(ctx, uid)
	checkoutLink := checkoutApiResponse.Data.Attributes.URL
	return checkoutLink, err
}

//// BindCustomer bind paid order to a customer.
//func (s *LemonSqueezyService) BindCustomer(ctx context.Context) error {
//	// 后端不相信前端送过来的产品选型和参数，而是选择自己通过order_id去查询 product & variant
//	// find order by order id, identifier
//	//orderApiResponse, response, err := s.Client.Orders.Get(ctx, "")
//	//orderApiResponse.Data.Attributes
//	// judge if order is bind already
//
//	// if bind already, return error
//
//	// if not bind, bind customer and user id.
//	return nil
//}
