package service

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/pkoukk/tiktoken-go"
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

func (s *OpenAIService) constructRequestCompletion(param *model.CompletionsReqBody) openai.ChatCompletionRequest {
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
	return req
}

// ChatCompletionStream chat completion
func (s *OpenAIService) ChatCompletionStream(ctx context.Context, param *model.CompletionsReqBody) (*openai.ChatCompletionStream, *openai.APIError) {
	req := s.constructRequestCompletion(param)
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
	req := s.constructRequestCompletion(param)
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
	// todo check if choices is empty.
	s.CountNonStreamTokens(param, []openai.ChatCompletionMessage{response.Choices[0].Message})
	return &response, nil
}

// Transcriptions transcription from audio to text
func (s *OpenAIService) Transcriptions(ctx context.Context, request *openai.AudioRequest) (openai.AudioResponse, error) {
	// todo save text
	return s.Client.CreateTranscription(ctx, *request)
}

// CountNonStreamTokens count message tokens and update uid quota used.
func (s *OpenAIService) CountNonStreamTokens(param *model.CompletionsReqBody, messages []openai.ChatCompletionMessage) {
	go func() {
		backgroundTokens := context.Background()
		req := s.constructRequestCompletion(param)
		messages = append(messages, req.Messages...)
		token, err := s.NumTokensFromMessages(messages, req.Model)
		if err != nil {
			log.Println("[CountNonStreamTokens]:NumTokensFromMessages err ", err)
		}
		if err = ChattyAIServiceInstance().UpdateTokenUsed(backgroundTokens, param.UID, int64(token)); err != nil {
			log.Println("[CountNonStreamTokens]: UpdateTokenUsed err ", err)
		}
	}()
}

// CountStreamTokens count message tokens and update uid quota used.
func (s *OpenAIService) CountStreamTokens(param *model.CompletionsReqBody, content string) {
	go func() {
		backgroundTokens := context.Background()
		req := s.constructRequestCompletion(param)
		promptToken, err := s.NumTokensFromMessages(req.Messages, req.Model)
		if err != nil {
			log.Println("[CountNonStreamTokens]:NumTokensFromMessages err ", err)
		}
		completionToken, err := s.GetTokenByModel(content, param.Model.ID)
		if err != nil {
			log.Println("[CountNonStreamTokens]:GetTokenByModel err ", err)
		}
		if err = ChattyAIServiceInstance().UpdateTokenUsed(backgroundTokens, param.UID, int64(promptToken+completionToken)); err != nil {
			log.Println("[CountNonStreamTokens]: UpdateTokenUsed err ", err)
		}
	}()
}

// GetTokenByModel get tokens by model.
func (s *OpenAIService) GetTokenByModel(text string, model string) (int, error) {
	// tiktoken
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		log.Println("[GetTokenByModel]: EncodingForModel err ", err)
		return 0, err
	}
	token := tkm.Encode(text, nil, nil)
	return len(token), nil
}

// NumTokensFromMessages get num tokens from input messages
func (s *OpenAIService) NumTokensFromMessages(messages []openai.ChatCompletionMessage, model string) (int, error) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("[NumTokensFromMessages]: EncodingForModel: %v", err)
		fmt.Println(err)
		return 0, err
	}

	var tokens_per_message int
	var tokens_per_name int
	if model == openai.GPT3Dot5Turbo0301 || model == openai.GPT3Dot5Turbo {
		tokens_per_message = 4
		tokens_per_name = -1
	} else if model == openai.GPT40314 || model == openai.GPT4 {
		tokens_per_message = 3
		tokens_per_name = 1
	} else {
		fmt.Println("[NumTokensFromMessages]: Warning: model not found. Using cl100k_base encoding.")
		tokens_per_message = 3
		tokens_per_name = 1
	}

	num_tokens := 0
	for _, message := range messages {
		num_tokens += tokens_per_message
		content := message.Name + message.Role + message.Content
		num_tokens += len(tkm.Encode(content, nil, nil))
		if message.Name != "" {
			num_tokens += tokens_per_name
		}
	}
	num_tokens += 3 // # every reply is primed with <|start|>assistant<|message|>
	return num_tokens, nil
}
