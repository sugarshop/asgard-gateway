package service

import (
	"context"
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
	if param.Meta.EventName == model.LemonSqueezyEventName_OrderCreated {
		// nil means order_created success
		return nil
	}
	return nil
}
