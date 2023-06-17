package service

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mitchellh/mapstructure"
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

// ListenWebhook Listen and deal with the lemon squeezy webhook request.
func (s *LemonSqueezyService) ListenWebhook(ctx context.Context, xSignature string, param *lemonsqueezy.WebhookRequest, rawBody []byte) error {
	// verify x-signature
	//if pass := remote.LemonSqueezyServiceInstance().Verify(ctx, xSignature, rawBody); !pass {
	//	err := fmt.Errorf("verify fail, xSignature %s param %+v", xSignature, param)
	//	log.Println("[ListenWebhook]: Verify err: ", err)
	//	return err
	//}
	// save order
	if param.Meta.EventName == string(model.LemonSqueezyEventName_OrderCreated) {
		// run in go-routine and background context.
		backgroundContext := context.Background()
		go func() {
			if err := LemonSqueezyServiceInstance().OrderCreatedEvent(backgroundContext, param); err != nil {
				log.Println("[ListenWebhook]: OrderCreatedEvent err: ", err)
			}
		}()
		// nil means order_created success
		return nil
	}
	if param.Meta.EventName == string(model.LemonSqueezyEventName_LicenseKeyCreated) {
		// nil means licenseKey_created success
		// todo save licenseKey record.
		return nil
	}

	err := fmt.Errorf("listen failed, event not found, %+v", param)
	log.Println(err)
	return err
}

// OrderCreatedEvent handle order created event.
func (s *LemonSqueezyService) OrderCreatedEvent(ctx context.Context, param *lemonsqueezy.WebhookRequest) error {
	orderAttr := lemonsqueezy.OrderAttributes{}
	convertedAttributes, ok := param.Data.Attributes.(map[string]interface{})
	if !ok {
		err := fmt.Errorf("attributes assertion failed")
		log.Println("[OrderCreatedEvent]: assertion fail", err)
		return err
	}
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &orderAttr,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	decoder.Decode(convertedAttributes)

	// incompletely convert from map to struct.
	createAt, ok := convertedAttributes["created_at"]
	if !ok {
		err := fmt.Errorf("attributes get created_at %v failed", createAt)
		log.Println("[OrderCreatedEvent]: fail", err)
		return err
	}
	createdAt, err := time.Parse(time.RFC3339Nano, fmt.Sprintf("%s", createAt))
	if err != nil {
		log.Println("[OrderCreatedEvent]: time.Parse failed", err)
	}
	order, err := dal.LemonSqueezyOrderDaoInstance().GetByOrderID(ctx, orderAttr.FirstOrderItem.OrderID)
	if err != nil {
		log.Println("[OrderCreatedEvent]: GetByOrderID err: ", err)
		return err
	}
	// found an order exsited already.
	if order.OrderID > 0 {
		err := fmt.Errorf("order of id: %d existed already. could not create a new order ", order.OrderID)
		log.Println("[OrderCreatedEvent]: err: ", err)
		return err
	}

	if param.Meta.CustomData == nil {
		err := fmt.Errorf("custom data nil")
		log.Println("[OrderCreatedEvent]: err: ", err)
		return err
	}
	uid, ok := param.Meta.CustomData["uid"]
	if !ok {
		err := fmt.Errorf("custom data parse uid fail, custom data: %+v", param.Meta.CustomData)
		log.Println("[OrderCreatedEvent]: err: ", err)
		return err
	}
	uidString := fmt.Sprintf("%v", uid)
	// create an order.
	if err := dal.LemonSqueezyOrderDaoInstance().Create(ctx, &model.LemonSqueezyOrder{
		OrderID:         orderAttr.FirstOrderItem.OrderID,
		UID:             uidString,
		StoreID:         orderAttr.StoreID,
		Identifier:      orderAttr.Identifier,
		Status:          orderAttr.Status,
		ProductID:       orderAttr.FirstOrderItem.ProductID,
		VariantID:       orderAttr.FirstOrderItem.VariantID,
		ProductName:     orderAttr.FirstOrderItem.ProductName,
		VariantName:     orderAttr.FirstOrderItem.VariantName,
		OrderCreateTime: createdAt,
	}); err != nil {
		log.Println("[OrderCreatedEvent]: Create err: ", err)
		return err
	}
	// todo create subscription strategy.
	if orderAttr.Status == model.LemonSqueezyOrderStatus_Paid {
		level := model.ChattyAIRightsLevel(orderAttr.FirstOrderItem.VariantName)
		// create or update subscription record.
		if err := ChattyAIServiceInstance().UpdateSubscription(ctx, uidString, level); err != nil {
			log.Println("[OrderCreatedEvent]: UpdateSubscription err: ", err)
			return err
		}
		return nil
	} else if orderAttr.Status == model.LemonSqueezyOrderStatus_Pending {
		// todo pending logic
		//return err
	} else if orderAttr.Status == model.LemonSqueezyOrderStatus_Refund {
		// todo DeleteSubscription
	} else if orderAttr.Status == model.LemonSqueezyOrderStatus_Failed {
		// todo deal with refund logic
	}
	return nil
}
