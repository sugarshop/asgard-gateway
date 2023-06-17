package service

import (
	"context"
	"log"
	"sync"

	"github.com/sugarshop/asgard-gateway/dal"
	"github.com/sugarshop/asgard-gateway/model"
)

// ChattyAIService a service that control chattyai subscription and bussiness logic.
type ChattyAIService struct {
}

var (
	chattyaiService *ChattyAIService
	chattyaiOnce    sync.Once
)

func ChattyAIServiceInstance() *ChattyAIService {
	chattyaiOnce.Do(func() {
		chattyaiService = &ChattyAIService{}
	})
	return chattyaiService
}

// CreateFreeSubscription create a ChattyAI subscription, when user first registered, create it.
func (s *ChattyAIService) CreateFreeSubscription(ctx context.Context) error {
	rights := &model.ChattyAIRights{}
	// init as free level rights
	rights.RenewalByLevel(model.ChattyAIRightsLevel_Free)
	if err := dal.ChattyAIRightsDaoInstance().Create(ctx, rights); err != nil {
		log.Println("[CreateFreeSubscription]: Create err: ", err)
		return err
	}
	return nil
}

// UpdateSubscription update a ChattyAI subscription
func (s *ChattyAIService) UpdateSubscription(ctx context.Context, uid string, level model.ChattyAIRightsLevel) error {
	if err := dal.ChattyAIRightsDaoInstance().UpdateLevel(ctx, uid, level); err != nil {
		log.Println("[UpdateSubscription]: UpdateLevel err ", err)
		return err
	}
	return nil
}
