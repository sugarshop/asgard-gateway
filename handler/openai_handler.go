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
	"time"
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
		return fmt.Errorf("StatusCode: %d, Type:%s, Code:%v", apiErr.HTTPStatusCode, apiErr.Type, apiErr.Code)
	}

	var curCompletion model.Completion
	ch := make(chan bool)

	gone := c.Stream(func(w io.Writer) bool {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Printf("[Completion]: Stream finished")
			go func() {
				if saveErr := curCompletion.Save(); saveErr != nil {
					fmt.Printf("[Completions]: failed to save completion: %+v", err)
					ch <- false
				} else {
					ch <- true
				}
				close(ch)
			}()
			return false
		}

		if err != nil {
			fmt.Printf("[Completion]: Stream error: %v\n", err)
			return false
		}
		// wrap completion.
		h.constructCompletion(&curCompletion, response)

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
		// client disconnected in middle of stream
		// do something after client is gone
		log.Println("client is gone")
	}

	select {
	case succ := <- ch:
		if succ {
			fmt.Printf("save success")
		} else {
			fmt.Printf("save failed")
		}
	case <-time.After(10 * time.Second): // 超时时间为10秒
		fmt.Println("save to db timeout!")
	}
	return nil
}

func (h *OpenAIHandler) constructCompletion(cur *model.Completion, input interface{}) {
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

	if len(cur.ChatID) == 0 {
		cur.ChatID = id
	}
	if len(cur.Model) == 0 {
		cur.Model = chatModel
	}
	if len(cur.Role) == 0 {
		cur.Role = role
	}
	cur.Content += content
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
		fmt.Printf("ChatCompletionStream error Code: %v, Type:%s, StatusCode: %d", apiErr.Code, apiErr.Type, apiErr.HTTPStatusCode)
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