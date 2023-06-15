package service

import (
	"context"
	"log"
	"reflect"
	"sync"

	"github.com/sugarshop/asgard-gateway/dal"
	"github.com/sugarshop/asgard-gateway/model"
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

// OrderCreatedEvent handle order created event.
func (s *LemonSqueezyService) OrderCreatedEvent(ctx context.Context, param *lemonsqueezy.WebhookRequest) error {
	orderAttr := &lemonsqueezy.OrderAttributes{}
	if err := convertMapToStruct(param.Data.Attributes, orderAttr); err != nil {
		log.Println("[OrderCreatedEvent]: convertMapToStruct err: ", err)
		return err
	}
	// fixme create or update ?
	if err := dal.LemonSqueezyOrderDaoInstance().Create(ctx, &model.LemonSqueezyOrder{
		OrderID:         orderAttr.FirstOrderItem.OrderID,
		UID:             "",
		StoreID:         orderAttr.StoreID,
		Identifier:      orderAttr.Identifier,
		Status:          orderAttr.Status,
		ProductID:       orderAttr.FirstOrderItem.ProductID,
		VariantID:       orderAttr.FirstOrderItem.VariantID,
		ProductName:     orderAttr.FirstOrderItem.ProductName,
		VariantName:     orderAttr.FirstOrderItem.VariantName,
		OrderCreateTime: orderAttr.CreatedAt,
	}); err != nil {
		log.Println("[OrderCreatedEvent]: Create err: ", err)
		return err
	}
	// todo create subscription strategy.
	if orderAttr.Status == model.LemonSqueezyOrderStatus_Paid {
		// create or update subscription record.
		return nil
	}
	return nil
}

func convertMapToStruct(m map[string]interface{}, s interface{}) error {
	stValue := reflect.ValueOf(s).Elem()
	sType := stValue.Type()
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		if value, ok := m[field.Name]; ok {
			stValue.Field(i).Set(reflect.ValueOf(value))
		}
	}
	return nil
}
