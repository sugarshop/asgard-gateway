package model

import (
	"time"
)

// ChattyAIRightsLevel level of chattyai rights
type ChattyAIRightsLevel string

const (
	ChattyAIRightsLevel_Free     ChattyAIRightsLevel = "Free"
	ChattyAIRightsLevel_Basic    ChattyAIRightsLevel = "Basic"
	ChattyAIRightsLevel_Advanced ChattyAIRightsLevel = "Advanced"
	ChattyAIRightsLevel_Pro      ChattyAIRightsLevel = "Pro"
)

type ChattyAIRights struct {
	ID                     uint64    `gorm:"column:id"`
	UID                    string    `gorm:"column:uid"`
	TokenQuota             int64     `gorm:"column:token_quota"`
	TokenUsed              int64     `gorm:"column:token_used"`
	TokenUsedTotal         int64     `gorm:"column:token_used_total"`
	ConversationQuota      int64     `gorm:"column:conversation_quota"`
	ConversationUsed       int64     `gorm:"column:conversation_used"`
	ConversationUsedTotal  int64     `gorm:"column:conversation_used_total"`
	AssistantQuota         int64     `gorm:"column:assistant_quota"`
	AssistantUsed          int64     `gorm:"column:assistant_used"`
	GPT4Access             bool      `gorm:"column:gpt4_access"`
	APIAccess              bool      `gorm:"column:api_access"`
	SubscriptionDate       time.Time `gorm:"column:subscription_date"` // subscription date means the first time when subscription occurred.
	SubscriptionUpdateDate time.Time `gorm:"column:subscription_update_date"`
	SubscriptionEndDate    time.Time `gorm:"column:subscription_end_date"`
	CreatedAt              time.Time `gorm:"column:created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at"`
}

// RenewalByLevel renewal rights or just init it.
func (s *ChattyAIRights) RenewalByLevel(level ChattyAIRightsLevel) {
	// clear usage data
	s.TokenUsed = 0
	s.ConversationUsed = 0
	s.AssistantUsed = 0

	// update level
	s.TokenQuota = GetTokenQuotaByLevel(level)
	s.ConversationQuota = GetConversationQuotaByLevel(level)
	s.AssistantQuota = GetAssistantQuotaByLevel(level)
	s.GPT4Access = GetGPT4Access(level)
	s.APIAccess = GetAPIAccess(level)

	// if subscription is not set yet, set now.
	if !s.SubscriptionDate.Before(time.Now()) {
		s.SubscriptionDate = time.Now()
	}
	// update subscription date
	s.SubscriptionUpdateDate = time.Now()
	s.SubscriptionEndDate = time.Now().AddDate(0, 1, 0)
}

func GetTokenQuotaByLevel(level ChattyAIRightsLevel) int64 {
	val := 100000
	switch level {
	case ChattyAIRightsLevel_Free:
		val = 100000
	case ChattyAIRightsLevel_Basic:
		val = 1000000
	case ChattyAIRightsLevel_Advanced:
		val = 2000000
	case ChattyAIRightsLevel_Pro:
		val = 5000000
	default:
		val = 100000
	}
	return int64(val)
}

func GetConversationQuotaByLevel(level ChattyAIRightsLevel) int64 {
	val := 500
	switch level {
	case ChattyAIRightsLevel_Free:
		val = 500
	case ChattyAIRightsLevel_Basic:
		val = 1200
	case ChattyAIRightsLevel_Advanced:
		val = 3000
	case ChattyAIRightsLevel_Pro:
		val = 8000
	default:
		val = 500
	}
	return int64(val)
}

func GetAssistantQuotaByLevel(level ChattyAIRightsLevel) int64 {
	val := 8
	switch level {
	case ChattyAIRightsLevel_Free:
		val = 8
	case ChattyAIRightsLevel_Basic:
		val = 15
	case ChattyAIRightsLevel_Advanced:
		val = 30
	case ChattyAIRightsLevel_Pro:
		val = 50
	default:
		val = 8
	}
	return int64(val)
}

func GetGPT4Access(level ChattyAIRightsLevel) bool {
	switch level {
	case ChattyAIRightsLevel_Free:
		return false
	case ChattyAIRightsLevel_Basic:
	case ChattyAIRightsLevel_Advanced:
	case ChattyAIRightsLevel_Pro:
		return true
	}
	return false
}

func GetAPIAccess(level ChattyAIRightsLevel) bool {
	switch level {
	case ChattyAIRightsLevel_Free:
		return false
	case ChattyAIRightsLevel_Basic:
	case ChattyAIRightsLevel_Advanced:
	case ChattyAIRightsLevel_Pro:
		return true
	}
	return false
}
