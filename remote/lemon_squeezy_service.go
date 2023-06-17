package remote

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/sugarshop/env"
	lemonsqueezy "github.com/sugarshop/lemonsqueezy-go"
)

// LemonSqueezyService lemon squeezy service
type LemonSqueezyService struct {
	Client              *lemonsqueezy.Client
	StoreID             string
	AssociatedVariantID string
}

var (
	lemonSqueezyService *LemonSqueezyService
	lemonSqueezyOnce    sync.Once
)

// LemonSqueezyServiceInstance lemon squeezy service instance
func LemonSqueezyServiceInstance() *LemonSqueezyService {
	storeID, ok := env.GlobalEnv().Get("LEMONSQUEEZYSTOREID")
	if !ok {
		log.Println("no LEMONSQUEEZYSTOREID env set")
	}
	variantID, ok := env.GlobalEnv().Get("LEMONSQUEEZYASSOCIATEDVARIANTID")
	if !ok {
		log.Println("no LEMONSQUEEZYASSOCIATEDVARIANTID env set")
	}
	apiKey, ok := env.GlobalEnv().Get("LEMONSQUEEZYAPIKEY")
	if !ok {
		log.Println("no LEMONSQUEEZYAPIKEY env set")
	}
	lemonSqueezyOnce.Do(func() {
		lemonSqueezyService = &LemonSqueezyService{
			Client:              lemonsqueezy.New(lemonsqueezy.WithAPIKey(apiKey)),
			StoreID:             storeID,
			AssociatedVariantID: variantID,
		}
	})
	return lemonSqueezyService
}

// ListCustomers get custormers list
func (s *LemonSqueezyService) ListCustomers(ctx context.Context) (*lemonsqueezy.CustomersApiResponse, error) {
	customerApiResponse, response, err := s.Client.Customers.List(ctx)
	if response.HTTPResponse.StatusCode != http.StatusOK {
		log.Println("[ListCustomers]: err:", err)
		return nil, err
	}
	return customerApiResponse, err
}

// ListVariants get custormers list
func (s *LemonSqueezyService) ListVariants(ctx context.Context) (*lemonsqueezy.VariantsApiResponse, error) {
	variantsApiResponse, response, err := s.Client.Variants.List(ctx)
	if response.HTTPResponse.StatusCode != http.StatusOK {
		log.Println("[ListVariants]: err:", err)
		return nil, err
	}
	return variantsApiResponse, err
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
		StoreID:         s.StoreID,
		VariantID:       s.AssociatedVariantID,
	})

	if response.HTTPResponse.StatusCode != http.StatusCreated {
		log.Println("[CreateCheckout]: err", err)
		return nil, err
	}
	return checkoutApiResponse, nil
}

// Verify verify if webhook is from lemonsqueezy
func (s *LemonSqueezyService) Verify(ctx context.Context, signature string, body []byte) bool {
	succ := s.Client.Webhooks.Verify(ctx, signature, body)
	log.Println("[Verify]: success ", succ)
	return succ
}
