package model

type LemonsqueezyEventname string

const (
	LemonSqueezyEventName_OrderCreated                 LemonsqueezyEventname = "order_created"
	LemonSqueezyEventName_OrderRefunded                LemonsqueezyEventname = "order_refunded"
	LemonSqueezyEventName_SubscriptionCreated          LemonsqueezyEventname = "subscription_created"
	LemonSqueezyEventName_SubscriptionUpdated          LemonsqueezyEventname = "subscription_updated"
	LemonSqueezyEventName_SubscriptionCancelled        LemonsqueezyEventname = "subscription_cancelled"
	LemonSqueezyEventName_SubscriptionResumed          LemonsqueezyEventname = "subscription_resumed"
	LemonSqueezyEventName_SubscriptionExpired          LemonsqueezyEventname = "subscription_expired"
	LemonSqueezyEventName_SubscriptionPaused           LemonsqueezyEventname = "subscription_paused"
	LemonSqueezyEventName_SubscriptionUnpaused         LemonsqueezyEventname = "subscription_unpaused"
	LemonSqueezyEventName_SubscriptionPaymentFailed    LemonsqueezyEventname = "subscription_payment_failed"
	LemonSqueezyEventName_SubscriptionPaymentSuccess   LemonsqueezyEventname = "subscription_payment_success"
	LemonSqueezyEventName_SubscriptionPaymentRecovered LemonsqueezyEventname = "subscription_payment_recovered"
	LemonSqueezyEventName_LicenseKeyCreated            LemonsqueezyEventname = "license_key_created"
)
