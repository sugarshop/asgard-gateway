package model

import (
	"time"
)

const (
	LemonSqueezyOrderStatus_Paid    = "paid"
	LemonSqueezyOrderStatus_Pending = "pending"
	LemonSqueezyOrderStatus_Failed  = "failed"
	LemonSqueezyOrderStatus_Refund  = "refunded"
)

type LemonSqueezyOrder struct {
	ID              int       `gorm:"column:id"`
	OrderID         int       `gorm:"column:order_id"`
	UID             string    `gorm:"column:uid"`
	StoreID         int       `gorm:"column:store_id"`
	Identifier      string    `gorm:"column:identifier"`
	Status          string    `gorm:"column:status"`
	ProductID       int       `gorm:"column:product_id"`
	VariantID       int       `gorm:"column:variant_id"`
	ProductName     string    `gorm:"column:product_name"`
	VariantName     string    `gorm:"column:variant_name"`
	OrderCreateTime time.Time `gorm:"column:order_create_time"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
