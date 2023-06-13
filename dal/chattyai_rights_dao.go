package dal

import (
	"context"
	"log"

	"github.com/sugarshop/asgard-gateway/db"
	"github.com/sugarshop/asgard-gateway/model"
)

type ChattyAIRightsDao struct {
}

const chattyaiRightsTable = "chattyai_rights"

var chattyaiRightsDao = &ChattyAIRightsDao{}

// ChattyAIRightsDaoInstance instance of ChattyAIRightsDao
func ChattyAIRightsDaoInstance() *ChattyAIRightsDao {
	return chattyaiRightsDao
}

// Create chattyai rights
func (d *ChattyAIRightsDao) Create(ctx context.Context, m *model.ChattyAIRights) error {
	if err := db.SugarShopDB().Table(chattyaiRightsTable).WithContext(ctx).Create(&m).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
