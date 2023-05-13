package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/sugarshop/asgard-gateway/model"
	"io"
	"log"
	"net/http"
)

type OpenAIHandler struct {
}

func NewOpenAIHandler() *OpenAIHandler {
	return &OpenAIHandler{}
}

func (h *OpenAIHandler) Register(e *gin.Engine)  {
	e.POST("/v1/openai/chat/completions", StreamWrapper(h.Completions))
}

func (h *OpenAIHandler) Completions(c *gin.Context) error {
	var reqBody model.CompletionsReqBody
	// bind json to reqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	stream, apiErr := h.OpenAIStream(c, &reqBody)
	if apiErr != nil {
		return fmt.Errorf("StatusCode: %d, Type:%s, Code:%s", apiErr.HTTPStatusCode, apiErr.Type, apiErr.Code)
	}

	gone := c.Stream(func(w io.Writer) bool {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Printf("[Completion]: Stream finished")
			return false
		}

		if err != nil {
			fmt.Printf("[Completion]: Stream error: %v\n", err)
			return false
		}

		jsonBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("[Completion]: err:%v", err)
			return false
		}
		formattedData := fmt.Sprintf("data: %s\n\n", string(jsonBytes))
		if _, err = w.Write([]byte(formattedData)); err != nil {
			fmt.Printf("[Completion]: Write data:%s, error: %v", formattedData, err)
			return false
		}
		return true
	})
	if gone {
		// do something after client is gone
		log.Println("client is gone")
	}
	return nil
}

// OpenAIStream return completion stream of the OpenAIChat
func (h *OpenAIHandler) OpenAIStream(c *gin.Context, param *model.CompletionsReqBody) (*openai.ChatCompletionStream, *openai.APIError) {
	client := openai.NewClient(param.Key)
	ctx := context.Background()
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	// fixme only for test env
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// default set to GPT3.5
	modelStr := openai.GPT3Dot5Turbo
	modelPtr := &param.Model
	if modelPtr != nil {
		modelStr = param.Model.ID
	}
	req := openai.ChatCompletionRequest{
		Model:     modelStr,
		MaxTokens: model.OPENAIMAXTOKENS,
		Messages: param.Messages,
		Stream: true,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		apiErr := buildError(err)
		fmt.Printf("ChatCompletionStream error Code: %s, Type:%s, StatusCode: %d", apiErr.Code, apiErr.Type, apiErr.HTTPStatusCode)
		return nil, apiErr
	}
	//defer stream.Close() // defer after call for stream
	return stream, nil
}

// buildErrorResp construct APIError
func buildError(err error) *openai.APIError {
	var aer *openai.APIError
	if errors.As(err, &aer) {
		return aer
	}
	// server deny
	aer.HTTPStatusCode = 500
	return aer
}