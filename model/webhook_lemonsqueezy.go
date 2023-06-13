package model

import "time"

// LemonSqueezyEventname event name for lemonsqueezy
type LemonSqueezyEventname string

const (
	LemonSqueezyEventName_OrderCreated                 LemonSqueezyEventname = "order_created"
	LemonSqueezyEventName_OrderRefunded                LemonSqueezyEventname = "order_refunded"
	LemonSqueezyEventName_SubscriptionCreated          LemonSqueezyEventname = "subscription_created"
	LemonSqueezyEventName_SubscriptionUpdated          LemonSqueezyEventname = "subscription_updated"
	LemonSqueezyEventName_SubscriptionCancelled        LemonSqueezyEventname = "subscription_cancelled"
	LemonSqueezyEventName_SubscriptionResumed          LemonSqueezyEventname = "subscription_resumed"
	LemonSqueezyEventName_SubscriptionExpired          LemonSqueezyEventname = "subscription_expired"
	LemonSqueezyEventName_SubscriptionPaused           LemonSqueezyEventname = "subscription_paused"
	LemonSqueezyEventName_SubscriptionUnpaused         LemonSqueezyEventname = "subscription_unpaused"
	LemonSqueezyEventName_SubscriptionPaymentFailed    LemonSqueezyEventname = "subscription_payment_failed"
	LemonSqueezyEventName_SubscriptionPaymentSuccess   LemonSqueezyEventname = "subscription_payment_success"
	LemonSqueezyEventName_SubscriptionPaymentRecovered LemonSqueezyEventname = "subscription_payment_recovered"
	LemonSqueezyEventName_LicenseKeyCreated            LemonSqueezyEventname = "license_key_created"
)

// LemonSqueezyRequest https://docs.lemonsqueezy.com/help/webhooks#event-types
type LemonSqueezyRequest struct {
	Data LemonSqueezyData `json:"data"`
	Meta LemonSqueezyMeta `json:"meta"`
}

// LemonSqueezyData request data of lemon squeezy webhook
type LemonSqueezyData struct {
	ID            string                 `json:"id"`    // have no idea but maybe some index?
	Type          string                 `json:"type"`  // orders、...
	Links         DataLinks              `json:"links"` //
	Attributes    map[string]interface{} `json:"attributes"`
	Relationships RelationshipData       `json:"relationships"`
}

// DataLinks DataLinks
type DataLinks struct {
	Self string `json:"self"`
}

// LemonSqueezyMeta meta infomation of Lemon Squeezy webhook
type LemonSqueezyMeta struct {
	TestMode  bool                  `json:"test_mode"`
	EventName LemonSqueezyEventname `json:"event_name"`
}

// RelationshipData RelationshipData
type RelationshipData struct {
	Store               RelationshipLinks `json:"store"`
	Customer            RelationshipLinks `json:"customer"`
	OrderItems          RelationshipLinks `json:"order-items"`
	LicenseKeys         RelationshipLinks `json:"license-keys"`
	Subscriptions       RelationshipLinks `json:"subscriptions"`
	DiscountRedemptions RelationshipLinks `json:"discount-redemptions"`
}

// RelationshipLinks RelationshipLinks
type RelationshipLinks struct {
	Links RelationshipLinksItem `json:"links"`
}

// RelationshipLinksItem RelationshipLinksItem
type RelationshipLinksItem struct {
	Self    string `json:"self"`
	Related string `json:"related"`
}

// OrderAttributes Order Attributes
type OrderAttributes struct {
	TestMode bool `json:"test_mode"`

	StoreID                int64             `json:"store_id"`                 // The ID of the store this order belongs to.
	CustomerID             int64             `json:"customer_id"`              // The ID of the customer this order belongs to.
	Identifier             string            `json:"identifier"`               // The unique identifier (UUID) for this order.
	OrderNumber            int64             `json:"order_number"`             // An integer representing the sequential order number for this store.
	UserName               string            `json:"user_name"`                // The full name of the customer.
	UserEmail              string            `json:"user_email"`               // The email address of the customer.
	Currency               string            `json:"currency"`                 // The ISO 4217 currency code for the order (e.g. USD, GBP, etc).
	CurrencyRate           string            `json:"currency_rate"`            // If the order currency is USD, this will always be 1.0. Otherwise, this is the currency conversion rate used to determine the cost of the order in USD at the time of purchase.
	Subtotal               int64             `json:"subtotal"`                 // A positive integer in cents representing the subtotal of the order in the order currency.
	DiscountTotal          int64             `json:"discount_total"`           // A positive integer in cents representing the total discount value applied to the order in the order currency.
	Tax                    float64           `json:"tax"`                      // A positive integer in cents representing the tax applied to the order in the order currency.
	Total                  int64             `json:"total"`                    // A positive integer in cents representing the total cost of the order in the order currency.
	SubtotalUsd            int64             `json:"subtotal_usd"`             // A positive integer in cents representing the subtotal of the order in USD.
	DiscountTotalUsd       int64             `json:"discount_total_usd"`       // A positive integer in cents representing the total discount value applied to the order in USD.
	TaxUsd                 float64           `json:"tax_usd"`                  // A positive integer in cents representing the tax applied to the order in USD.
	TotalUsd               int64             `json:"total_usd"`                // A positive integer in cents representing the total cost of the order in USD.
	TaxName                string            `json:"tax_name"`                 // If tax is applied to the order, this will be the name of the tax rate (e.g. VAT, Sales Tax, etc).
	TaxRate                string            `json:"tax_rate"`                 // If tax is applied to the order, this will be the rate of tax as a decimal percentage.
	Status                 string            `json:"status"`                   // The status of the order. One of pending, failed, paid, refunded.
	StatusFormatted        string            `json:"status_formatted"`         // The formatted status of the order.
	Refunded               bool              `json:"refunded"`                 // Has the value true if the order has been refunded.
	RefundedAt             time.Time         `json:"refunded_at"`              // If the order has been refunded, this will be an ISO-8601 formatted date-time string indicating when the order was refunded.
	SubtotalFormatted      string            `json:"subtotal_formatted"`       // A human-readable string representing the subtotal of the order in the order currency (e.g. $9.99).
	DiscountTotalFormatted string            `json:"discount_total_formatted"` // A human-readable string representing the total discount value applied to the order in the order currency (e.g. $9.99).
	TaxFormatted           string            `json:"tax_formatted"`            // A human-readable string representing the tax applied to the order in the order currency (e.g. $9.99).
	TotalFormatted         string            `json:"total_formatted"`          // A human-readable string representing the total cost of the order in the order currency (e.g. $9.99).
	FirstOrderItem         FirstOrderItem    `json:"first_order_item"`         // see more in
	Urls                   map[string]string `json:"urls"`                     // An object of customer-facing URLs for this order. It contains: receipt - A pre-signed URL for viewing the order in the customer's My Orders page.
	CreatedAt              time.Time         `json:"created_at"`               // An ISO-8601 formatted date-time string indicating when the object was created.
	UpdatedAt              time.Time         `json:"updated_at"`               // An ISO-8601 formatted date-time string indicating when the object was last updated.
}

// FirstOrderItem First Order Item
type FirstOrderItem struct {
	TestMode bool `json:"test_mode"`

	Price       int       `json:"price"`
	OrderID     int       `json:"order_id"`
	ProductID   int       `json:"product_id"`
	VariantID   int       `json:"variant_id"`
	ProductName string    `json:"product_name"`
	VariantName string    `json:"variant_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
