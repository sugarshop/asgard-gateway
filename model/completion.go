package model

import (
	"log"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/sugarshop/asgard-gateway/db"
)

type Completion struct {
	ID        int64     `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ChatID    string    `json:"chat_id"`
	Model     string    `json:"model"`
	Content   string    `json:"content"`
	Role      string    `json:"role"`
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

func (s *Completion) ConstructMessage(input interface{}) {
	streamResponse, streamOk := input.(openai.ChatCompletionStreamResponse)
	completionResponse, completionOk := input.(openai.ChatCompletionResponse)
	var id string
	var chatModel string
	var role string
	var content string
	if streamOk {
		id = streamResponse.ID
		chatModel = streamResponse.Model
		if len(streamResponse.Choices) > 0 {
			role = streamResponse.Choices[0].Delta.Role
		}
		content = streamResponse.Choices[0].Delta.Content
	}
	if completionOk {
		id = completionResponse.ID
		chatModel = completionResponse.Model
		if len(completionResponse.Choices) > 0 {
			role = completionResponse.Choices[0].Message.Role
		}
		content = completionResponse.Choices[0].Message.Content
	}

	if len(s.ChatID) == 0 {
		s.ChatID = id
	}
	if len(s.Model) == 0 {
		s.Model = chatModel
	}
	if len(s.Role) == 0 {
		s.Role = role
	}
	s.Content += content
}
