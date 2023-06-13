package model

import "time"

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
	SubscriptionDate       time.Time `gorm:"column:subscription_date"`
	SubscriptionUpdateDate time.Time `gorm:"column:subscription_update_date"`
	SubscriptionEndDate    time.Time `gorm:"column:subscription_end_date"`
	CreatedAt              time.Time `gorm:"column:created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at"`
}
