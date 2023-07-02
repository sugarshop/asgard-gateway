package model

import (
	"fmt"

	"github.com/sashabaranov/go-openai"
)

const (
	OPENAIMAXTOKENS = 1000
)

type OpenAIModel struct {
	ID string `json:"id"`
}

type ChatRequest struct {
	Model       string                         `json:"model,omitempty"`
	Messages    []openai.ChatCompletionMessage `json:"messages"`
	Temperature float64                        `json:"temperature"`
	Stream      bool                           `json:"stream"`
}

type OpenAIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}

func (e *OpenAIError) Error() string {
	return fmt.Sprintf("OpenAI error: %s (%s %s %s)", e.Message, e.Type, e.Param, e.Code)
}

type CompletionsReqBody struct {
	Model       OpenAIModel                    `json:"model"`
	Messages    []openai.ChatCompletionMessage `json:"messages"`
	Temperature float64                        `json:"temperature"`
	Stream      bool                           `json:"stream"`
	UID         string                         `json:"uid"`
}

// LogprobResult represents logprob result of Choice.
type LogprobResult struct {
	Tokens        []string             `json:"tokens"`
	TokenLogprobs []float32            `json:"token_logprobs"`
	TopLogprobs   []map[string]float32 `json:"top_logprobs"`
	TextOffset    []int                `json:"text_offset"`
}

// Usage Represents the total token usage per request to OpenAI.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// CompletionChoice represents one of possible completions.
type CompletionChoice struct {
	Text         string        `json:"text"`
	Index        int           `json:"index"`
	FinishReason string        `json:"finish_reason"`
	LogProbs     LogprobResult `json:"logprobs"`
}

// CompletionResponse represents a response structure for completion API.
type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int64              `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
	Usage   Usage              `json:"usage"`
}
