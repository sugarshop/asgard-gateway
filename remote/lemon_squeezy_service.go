package remote

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/sugarshop/asgard-gateway/model"
	lemonsqueezy "github.com/sugarshop/lemonsqueezy-go"
)

// LemonSqueezyService lemon squeezy service
type LemonSqueezyService struct {
	Client *lemonsqueezy.Client
}

var (
	lemonSqueezyService *LemonSqueezyService
	lemonSqueezyOnce    sync.Once
)

// LemonSqueezyServiceInstance lemon squeezy service instance
func LemonSqueezyServiceInstance() *LemonSqueezyService {
	lemonSqueezyOnce.Do(func() {
		lemonSqueezyService = &LemonSqueezyService{
			Client: lemonsqueezy.New(lemonsqueezy.WithAPIKey(model.LemonSqueezyAPIKey)),
		}
	})
	return lemonSqueezyService
}

// ListCustomers get custormers list
func (s *LemonSqueezyService) ListCustomers(ctx context.Context) (*lemonsqueezy.CustomersApiResponse, error) {
	customerApiResponse, response, err := s.Client.Customers.List(ctx)
	if response.HTTPResponse.StatusCode != http.StatusOK {
		log.Println("[ListCustomers]: err:%+v")
		return nil, err
	}
	return customerApiResponse, err
}

// CreateCheckout create checkout link
func (s *LemonSqueezyService) CreateCheckout(ctx context.Context, uid string) (*lemonsqueezy.CheckoutApiResponse, error) {
	expireDate := time.Now().AddDate(0, 0, 1)

	data := map[string]string{
		"uid": uid,
	}

	checkoutApiResponse, response, err := s.Client.Checkouts.Create(ctx, &lemonsqueezy.CheckoutCreateParams{
		EnabledVariants: []int{},
		ButtonColor:     "#2DD272",
		Embed:           true,
		Media:           false,
		Logo:            true,
		CustomData:      data,
		ExpiresAt:       expireDate,
		StoreID:         model.LemonSqueezy_StoreID,
		VariantID:       model.LemonSqueezy_Associated_VariantID,
	})

	if response.HTTPResponse.StatusCode != http.StatusOK {
		log.Println("[CreateCheckout]: err:%+v")
		return nil, err
	}
	return checkoutApiResponse, nil
}
