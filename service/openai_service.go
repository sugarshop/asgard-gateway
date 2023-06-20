package service

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/sashabaranov/go-openai"
	"github.com/sugarshop/asgard-gateway/model"
	"github.com/sugarshop/env"
)

// OpenAIService a service that connect openai service
type OpenAIService struct {
	Client *openai.Client
}

var (
	openaiService *OpenAIService
	openaiOnce    sync.Once
)

func OpenAIServiceInstance() *OpenAIService {
	apiKey, ok := env.GlobalEnv().Get("OPENAIAPIKEY")
	if !ok {
		log.Println("no OPENAIAPIKEY env set")
		// todo alert err
	}
	openaiOnce.Do(func() {
		openaiService = &OpenAIService{
			Client: openai.NewClient(apiKey),
		}
	})
	return openaiService
}

// ChatCompletionStream chat completino
func (s *OpenAIService) ChatCompletionStream(ctx context.Context, param *model.CompletionsReqBody) (*openai.ChatCompletionStream, *openai.APIError) {
	// default set to GPT3.5
	modelStr := openai.GPT3Dot5Turbo
	modelPtr := &param.Model
	if modelPtr != nil {
		modelStr = param.Model.ID
	}
	message := param.Messages
	if len(param.Messages) >= 5 {
		message = append([]openai.ChatCompletionMessage{param.Messages[0]}, param.Messages[len(param.Messages)-3:]...)
	}
	log.Printf("model: %v message length:%d", param.Model, len(param.Messages))
	req := openai.ChatCompletionRequest{
		Model:     modelStr,
		MaxTokens: model.OPENAIMAXTOKENS,
		Messages:  message,
	}
	stream, err := s.Client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		apiErr := &openai.APIError{
			Code:           model.ErrCodeInternalServerError,
			Message:        err.Error(),
			HTTPStatusCode: 500,
		}
		fmt.Printf("ChatCompletionStream error Code: %v, Type:%s, StatusCode: %d", apiErr.Code, apiErr.Type, apiErr.HTTPStatusCode)
		return nil, apiErr
	}
	return stream, nil
}

// ChatCompletion chat completino
func (s *OpenAIService) ChatCompletion(ctx context.Context, param *model.CompletionsReqBody) (*openai.ChatCompletionResponse, *openai.APIError) {
	// default set to GPT3.5
	modelStr := openai.GPT3Dot5Turbo
	modelPtr := &param.Model
	if modelPtr != nil {
		modelStr = param.Model.ID
	}
	message := param.Messages
	if len(param.Messages) >= 5 {
		message = append([]openai.ChatCompletionMessage{param.Messages[0]}, param.Messages[len(param.Messages)-3:]...)
	}
	log.Printf("model: %v message length:%d", param.Model, len(param.Messages))
	req := openai.ChatCompletionRequest{
		Model:     modelStr,
		MaxTokens: model.OPENAIMAXTOKENS,
		Messages:  message,
	}
	response, err := s.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		apiErr := &openai.APIError{
			Code:           model.ErrCodeInternalServerError,
			Message:        err.Error(),
			HTTPStatusCode: 500,
		}
		fmt.Printf("ChatCompletionStream error Code: %v, Type:%s, StatusCode: %d", apiErr.Code, apiErr.Type, apiErr.HTTPStatusCode)
		return nil, apiErr
	}
	return &response, nil
}

// Transcriptions transcription from audio to text
func (s *OpenAIService) Transcriptions(ctx context.Context, request *openai.AudioRequest) (openai.AudioResponse, error) {
	// todo save text
	return s.Client.CreateTranscription(ctx, *request)
}
