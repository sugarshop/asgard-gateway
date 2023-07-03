package service

import (
	"context"
	"fmt"
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

// GetSubscriptionByUID get subscription by uid
func (s *ChattyAIService) GetSubscriptionByUID(ctx context.Context, uid string) (*model.ChattyAIRights, error) {
	rights, err := dal.ChattyAIRightsDaoInstance().GetByUID(ctx, uid)
	if err != nil {
		log.Println("[GetSubscriptionByUID]: GetByUID err: ", err)
		return nil, err
	}
	return rights, nil
}

// TokenIsSufficient chatty_ai subscription token is sufficient.
func (s *ChattyAIService) TokenIsSufficient(ctx context.Context, uid string) (bool, error) {
	if len(uid) <= 0 {
		log.Println("empty uid")
		return false, fmt.Errorf("empty uid")
	}
	rights, err := dal.ChattyAIRightsDaoInstance().GetByUID(ctx, uid)
	if err != nil {
		log.Println("[TokenIsSufficient]: GetByUID err: ", err)
		return false, err
	}
	return rights.TokenIsSufficient(), nil
}

// APIAccess weather has api access
func (s *ChattyAIService) APIAccess(ctx context.Context, apikey string) (bool, error) {
	// using uid as api_key
	if len(apikey) <= 0 {
		log.Println("empty api_key")
		return false, fmt.Errorf("empty api_key")
	}
	uid := apikey
	rights, err := dal.ChattyAIRightsDaoInstance().GetByUID(ctx, uid)
	if err != nil {
		log.Println("[APIAccess]: GetByUID err: ", err)
		return false, err
	}
	return rights.APIAccess && rights.TokenIsSufficient(), nil
}

// GPT4Access weather has gpt4 access
func (s *ChattyAIService) GPT4Access(ctx context.Context, uid string) (bool, error) {
	rights, err := dal.ChattyAIRightsDaoInstance().GetByUID(ctx, uid)
	if err != nil {
		log.Println("[GPT4Access]: GetByUID err: ", err)
		return false, err
	}
	return rights.GPT4Access, nil
}

// AssistantIsSufficient chatty_ai subscription assistant is sufficient.
func (s *ChattyAIService) AssistantIsSufficient(ctx context.Context, uid string) (bool, error) {
	rights, err := dal.ChattyAIRightsDaoInstance().GetByUID(ctx, uid)
	if err != nil {
		log.Println("[AssistantIsSufficient]: GetByUID err: ", err)
		return false, err
	}
	return rights.AssistantIsSuuffcient(), nil
}

// CreateFreeSubscription create a ChattyAI subscription, when user first registered, create it.
func (s *ChattyAIService) CreateFreeSubscription(ctx context.Context, uid string) error {
	rights := &model.ChattyAIRights{
		UID: uid,
	}
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

// UpdateTokenUsed update user's token used service
func (s *ChattyAIService) UpdateTokenUsed(ctx context.Context, uid string, token int64) error {
	if err := dal.ChattyAIRightsDaoInstance().UpdateTokenUsed(ctx, uid, token); err != nil {
		log.Println("[UpdateTokenUsed]: UpdateTokenUsed err ", err)
		return err
	}
	return nil
}
