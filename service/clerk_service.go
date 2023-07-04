package service

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/sugarshop/asgard-gateway/model"
	"github.com/sugarshop/env"
	svix "github.com/svix/svix-webhooks/go"
)

// ClerkService a service deal with clerk login event.
type ClerkService struct {
	Wh *svix.Webhook
}

var (
	clerkService *ClerkService
	clerkOnce    sync.Once
)

// ClerkServiceInstance add clerk service
func ClerkServiceInstance() *ClerkService {
	secret, ok := env.GlobalEnv().Get("CLERKSECRET")
	if !ok {
		log.Println("[ClerkServiceInstance]: clerksecret get failed")
	}
	clerkOnce.Do(func() {
		wh, err := svix.NewWebhook(secret)
		if err != nil {
			log.Println("[ClerkServiceInstance]: svix.NewWebhook err ", err)
		}
		clerkService = &ClerkService{
			Wh: wh,
		}
	})
	return clerkService
}

func (s *ClerkService) ListenClerkWebHook(ctx context.Context, event *model.ClerkEvent, payload []byte, header http.Header) error {
	// verify webhook from endpoint
	if err := s.Wh.Verify(payload, header); err != nil {
		log.Println("[ListenClerkWebHook]: Verify err ", err)
		return err
	}
	if event.Type == model.ClerkWebhookEvent_USERCREATED {
		go func() {
			backgroundCtx := context.Background()
			if err := ChattyAIServiceInstance().CreateFreeSubscription(backgroundCtx, event.Data.ID); err != nil {
				log.Println("[ListenClerkWebHook]: CreateFreeSubscription err", err)
			}
		}()
	} else {
		// todo other event
		log.Println("[ListenClerkWebHook]: other event ", event.Type, payload)
	}
	return nil
}
