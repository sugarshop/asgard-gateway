package dal

import (
	"context"
	"fmt"
	"log"

	"github.com/sugarshop/asgard-gateway/db"
	"github.com/sugarshop/asgard-gateway/model"
	"gorm.io/gorm"
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
	rights, err := d.TxGetByUID(ctx, db.SugarShopDB(), uid)
	if err != nil {
		log.Println("[ChattyAIRightsDao]: GetByUID err ", err)
		return nil, err
	}
	return rights, nil
}

func (d *ChattyAIRightsDao) TxGetByUID(ctx context.Context, tx *gorm.DB, uid string) (*model.ChattyAIRights, error) {
	rights := &model.ChattyAIRights{}
	err := tx.WithContext(ctx).Table(chattyaiRightsTable).Where("UID = ?", uid).Take(rights).Error
	if err != nil {
		log.Println("[ChattyAIRightsDao]: TxGetByUID err ", err)
		return nil, err
	}
	return rights, nil
}

// UpdateLevel chattyai rights update
func (d *ChattyAIRightsDao) UpdateLevel(ctx context.Context, uid string, level model.ChattyAIRightsLevel) error {
	log.Println("[UpdateLevel]: UpdateLevel ", level)
	err := db.SugarShopDB().WithContext(ctx).Table(chattyaiRightsTable).Transaction(func(tx *gorm.DB) error {
		right, txErr := d.TxGetByUID(ctx, tx, uid)
		if txErr != nil {
			log.Println("[UpdateLevel]: TxGetByUID err ", txErr)
			return txErr
		}
		// renewal it
		right.RenewalByLevel(level)
		session := tx.Save(right)
		if session.Error != nil {
			log.Println("[UpdateLevel]: Save err ", session.Error)
			return session.Error
		}
		if session.RowsAffected == 0 {
			return fmt.Errorf("save failed, rowsaffected = 0, uid %s", uid)
		}
		return nil
	})
	if err != nil {
		log.Println(ctx, "[UpdateLevel]: transaction failed. err ", err)
		return err
	}
	return nil
}

// UpdateTokenUsed update token used. use transaction to avoid concurrent r/w on the same line.
func (d *ChattyAIRightsDao) UpdateTokenUsed(ctx context.Context, uid string, token int64) error {
	log.Println("[UpdateTokenUsed]: UpdateTokenUsed ", token)
	err := db.SugarShopDB().WithContext(ctx).Table(chattyaiRightsTable).Transaction(func(tx *gorm.DB) error {
		right, txErr := d.TxGetByUID(ctx, tx, uid)
		if txErr != nil {
			log.Println("[UpdateTokenUsed]: TxGetByUID err ", txErr)
			return txErr
		}
		// update token used
		right.TokenUsed += token
		right.TokenUsedTotal += token
		session := tx.Exec(fmt.Sprintf("update %s set token_used = ?, token_used_total = ? where uid = ?", chattyaiRightsTable), right.TokenUsed, right.TokenUsedTotal, uid)
		if session.Error != nil {
			log.Println("[UpdateTokenUsed]: update err ", session.Error)
			return session.Error
		}
		if session.RowsAffected == 0 {
			return fmt.Errorf("update failed, rowsaffected = 0, uid %s", uid)
		}
		return nil
	})
	if err != nil {
		log.Println(ctx, "[UpdateTokenUsed]: transaction failed. err ", err)
		return err
	}
	return nil
}

// UpdateAssistantUsed update assistant used. use transaction to avoid concurrent r/w on the same line.
func (d *ChattyAIRightsDao) UpdateAssistantUsed(ctx context.Context, num int64, uid string) error {
	log.Println("[UpdateAssistantUsed]: UpdateAssistantUsed num ", num, "uid ", uid)
	err := db.SugarShopDB().WithContext(ctx).Table(chattyaiRightsTable).Transaction(func(tx *gorm.DB) error {
		right, txErr := d.TxGetByUID(ctx, tx, uid)
		if txErr != nil {
			log.Println("[UpdateAssistantUsed]: TxGetByUID err ", txErr)
			return txErr
		}
		// update assistant used
		right.AssistantUsed += num
		session := tx.Exec(fmt.Sprintf("update %s set assistant_used = ? where uid = ?", chattyaiRightsTable), right.AssistantUsed, uid)
		if session.Error != nil {
			log.Println("[UpdateAssistantUsed]: update err ", session.Error)
			return session.Error
		}
		if session.RowsAffected == 0 {
			return fmt.Errorf("update failed, rowsaffected = 0, uid %s", uid)
		}
		return nil
	})
	if err != nil {
		log.Println(ctx, "[UpdateAssistantUsed]: transaction failed. err ", err)
		return err
	}
	return nil
}
