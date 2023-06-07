package service

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/sugarshop/asgard-gateway/model"
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
func (s *WebHookService) ListenLemonSqueezy(ctx context.Context, param *model.LemonSqueezyRequest) error {
	// todo verfify test_mode
	if param.Meta.EventName == model.LemonSqueezyEventName_OrderCreated {
		// nil means order_created success
		return nil
	}
	if param.Meta.EventName == model.LemonSqueezyEventName_LicenseKeyCreated {
		// nil means licenseKey_created success
		return nil
	}
	err := fmt.Errorf("listen failed, event not found, %s", param)
	log.Println(err)
	return err
}
