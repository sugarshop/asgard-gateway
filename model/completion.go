package model

import (
	"github.com/sugarshop/asgard-gateway/db"
	"log"
	"time"
)

type Completion struct {
	ID        int64     `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ChatID    string    `json:"chat_id"`
	Model     string    `json:"model"`
	Content   string    `json:"content"`
	Role   	  string    `json:"role"`
}

func (s *Completion) Save() error {
	// set Now Time()
	s.CreatedAt = time.Now()
	var results []Completion
	err := db.CompletionDB().From("completion").Insert(s).Execute(&results)
	if err != nil {
		log.Printf("[Save]: err :%+v", err)
		return err
	}
	return nil
}