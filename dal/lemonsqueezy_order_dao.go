package dal

import (
	"context"
	"log"

	"github.com/sugarshop/asgard-gateway/db"
	"github.com/sugarshop/asgard-gateway/model"
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
		log.Println(err)
		return err
	}
	return nil
}
