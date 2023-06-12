package model

import "time"

type LemonSqueezyOrder struct {
	ID              uint64    `gorm:"column:id"`
	OrderID         uint64    `gorm:"column:order_id"`
	UID             string    `gorm:"column:uid"`
	StoreID         uint64    `gorm:"column:store_id"`
	Identifier      string    `gorm:"column:identifier"`
	Status          string    `gorm:"column:status"`
	ProductID       uint64    `gorm:"column:product_id"`
	VariantID       uint64    `gorm:"column:variant_id"`
	ProductName     string    `gorm:"column:product_name"`
	VariantName     string    `gorm:"column:variant_name"`
	OrderCreateTime time.Time `gorm:"column:order_create_time"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
