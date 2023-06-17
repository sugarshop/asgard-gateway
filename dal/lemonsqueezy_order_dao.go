package dal

import (
	"context"
	"log"

	"github.com/sugarshop/asgard-gateway/db"
	"github.com/sugarshop/asgard-gateway/model"
	"gorm.io/gorm"
)

type LemonSqueezyOrderDao struct {
}

const lemonsqueezyOrderTable = "lemonsqueezy_orders"

var lemonSqueezyOrderDao = &LemonSqueezyOrderDao{}

// LemonSqueezyOrderDaoInstance instance of LemonSqueezyOrderDao
func LemonSqueezyOrderDaoInstance() *LemonSqueezyOrderDao {
	return lemonSqueezyOrderDao
}

// Create lemonsqueezy order
func (d *LemonSqueezyOrderDao) Create(ctx context.Context, m *model.LemonSqueezyOrder) error {
	if err := db.SugarShopDB().Table(lemonsqueezyOrderTable).WithContext(ctx).Create(&m).Error; err != nil {
		log.Println("[LemonSqueezyOrderDao]: Create err ", err)
		return err
	}
	return nil
}

// GetByOrderID get by id.
func (d *LemonSqueezyOrderDao) GetByOrderID(ctx context.Context, orderID int) (*model.LemonSqueezyOrder, error) {
	order := &model.LemonSqueezyOrder{}
	err := db.SugarShopDB().Table(lemonsqueezyOrderTable).WithContext(ctx).Where("order_id = ?", orderID).Take(order).Error
	if (err != nil) && (err != gorm.ErrRecordNotFound) {
		log.Println("[LemonSqueezyOrderDao:GetByOrderID]: err: ", err)
		return nil, err
	}
	return order, nil
}
