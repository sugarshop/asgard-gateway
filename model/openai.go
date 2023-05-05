package model

import "fmt"

const (
	OPENAIAPIHOST = "https://api.openai.com"
	DEFAULTTEMPERATURE = 1
	NEXTPUBLICDEFAULTSYSTEMPROMPT = "You are ChatGPT, a large language model trained by OpenAI. Follow the user's instructions carefully. Respond using markdown."
	OPENAIAPITYPE = "openai"
	OPENAIAPIVERSION = "2023-03-15-preview"
	OPENAIORGANIZATION = ""
	AZUREDEPLOYMENTID = ""
	OPENAIAPIKEY = "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"
)

type OpenAIModel struct {
	ID string `json:"id"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model      string    `json:"model,omitempty"`
	Messages   []Message `json:"messages"`
	MaxTokens  int       `json:"max_tokens,omitempty"`
	Temperature float64  `json:"temperature"`
	Stream     bool      `json:"stream"`
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
	Model        OpenAIModel `json:"model"`
	SystemPrompt string            `json:"systemPrompt"`
	Temperature  float64           `json:"temperature"`
	Key          string            `json:"key"`
	Messages     []Message  `json:"messages"`
}