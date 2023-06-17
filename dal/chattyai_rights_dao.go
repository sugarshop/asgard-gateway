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
	if err := db.SugarShopDB().WithContext(ctx).Table(chattyaiRightsTable).Create(&m).Error; err != nil {
		log.Println("[ChattyAIRightsDao]: Create err ", err)
		return err
	}
	return nil
}

// GetByUID get by uid
func (d *ChattyAIRightsDao) GetByUID(ctx context.Context, uid string) (*model.ChattyAIRights, error) {
	rights := &model.ChattyAIRights{}
	err := db.SugarShopDB().WithContext(ctx).Table(chattyaiRightsTable).Where("UID = ?", uid).Take(rights).Error
	if err != nil {
		log.Println("[ChattyAIRightsDao]: GetByUID err ", err)
		return nil, err
	}
	return rights, nil
}

// UpdateLevel chattyai rights update
func (d *ChattyAIRightsDao) UpdateLevel(ctx context.Context, uid string, level model.ChattyAIRightsLevel) error {
	log.Println("[ChattyAIRightsDao]: UpdateLevel ", level)
	// find right first
	right, err := d.GetByUID(ctx, uid)
	if err != nil {
		log.Println("[ChattyAIRightsDao]: GetByUID err ", err)
		return err
	}
	// renewal it
	right.RenewalByLevel(level)
	// save it
	if err = db.SugarShopDB().WithContext(ctx).Table(chattyaiRightsTable).Save(&right).Error; err != nil {
		log.Println("[ChattyAIRightsDao]: Save err ", err)
		return err
	}
	return nil
}
