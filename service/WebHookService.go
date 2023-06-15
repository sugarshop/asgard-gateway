package service

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/sugarshop/asgard-gateway/model"
	"github.com/sugarshop/asgard-gateway/remote"
	lemonsqueezy "github.com/sugarshop/lemonsqueezy-go"
)

type WebHookService struct {
}

var (
	webhookService *WebHookService
	webhookOnce    sync.Once
)

func WebHookServiceInstance() *WebHookService {
	webhookOnce.Do(func() {
		webhookService = &WebHookService{}
	})
	return webhookService
}

// ListenLemonSqueezy Listen and deal with the lemon squeezy webhook request.
func (s *WebHookService) ListenLemonSqueezy(ctx context.Context, xSignature string, param *lemonsqueezy.WebhookRequest, rawBody []byte) error {
	// verify x-signature
	if err := remote.LemonSqueezyServiceInstance().Verify(ctx, xSignature, rawBody); err != nil {
		log.Println("[ListenLemonSqueezy]: Verify err: ", err)
		return err
	}
	// save order
	if param.Meta.EventName == string(model.LemonSqueezyEventName_OrderCreated) {
		if err := LemonSqueezyServiceInstance().OrderCreatedEvent(ctx, param); err != nil {
			log.Println("[ListenLemonSqueezy]: OrderCreatedEvent err: ", err)
			return err
		}
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
