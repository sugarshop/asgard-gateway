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
	ID                     uint64    `gorm:"column:id" json:"-"`
	UID                    string    `gorm:"column:uid" json:"-"`
	TokenQuota             int64     `gorm:"column:token_quota" json:"token_quota"`
	TokenUsed              int64     `gorm:"column:token_used" json:"token_used"`
	TokenUsedTotal         int64     `gorm:"column:token_used_total" json:"-"`
	AssistantQuota         int64     `gorm:"column:assistant_quota" json:"assistant_quota"`
	AssistantUsed          int64     `gorm:"column:assistant_used" json:"assistant_used"`
	GPT4Access             bool      `gorm:"column:gpt_4_access" json:"gpt_4_access"`
	APIAccess              bool      `gorm:"column:api_access" json:"api_access"`
	SubscriptionDate       time.Time `gorm:"column:subscription_date" json:"subscription_date"` // subscription date means the first time when subscription occurred.
	SubscriptionUpdateDate time.Time `gorm:"column:subscription_update_date" json:"subscription_update_date"`
	SubscriptionEndDate    time.Time `gorm:"column:subscription_end_date" json:"subscription_end_date"`
	CreatedAt              time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt              time.Time `gorm:"column:updated_at" json:"-"`
}

// TokenIsSufficient return if token is sufficient
func (s *ChattyAIRights) TokenIsSufficient() bool {
	return s.TokenQuota > s.TokenUsed
}

func (s *ChattyAIRights) AssistantIsSuuffcient() bool {
	return s.AssistantQuota > s.AssistantUsed
}

// RenewalByLevel renewal rights or just init it.
func (s *ChattyAIRights) RenewalByLevel(level ChattyAIRightsLevel) {
	// clear usage data
	s.TokenUsed = 0

	// update level
	s.TokenQuota = GetTokenQuotaByLevel(level)
	s.AssistantQuota = GetAssistantQuotaByLevel(level)
	s.GPT4Access = GetGPT4Access(level)
	s.APIAccess = GetAPIAccess(level)

	// if subscription is not set yet, set now.
	if s.SubscriptionDate.Before(time.Now()) {
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
	access := false
	switch level {
	case ChattyAIRightsLevel_Free:
		access = false
	case ChattyAIRightsLevel_Basic:
		access = true
	case ChattyAIRightsLevel_Advanced:
		access = true
	case ChattyAIRightsLevel_Pro:
		access = true

	}
	return access
}

func GetAPIAccess(level ChattyAIRightsLevel) bool {
	access := false
	switch level {
	case ChattyAIRightsLevel_Free:
		access = false
	case ChattyAIRightsLevel_Basic:
		access = true
	case ChattyAIRightsLevel_Advanced:
		access = true
	case ChattyAIRightsLevel_Pro:
		access = true
	}
	return access
}
