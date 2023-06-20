package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/model"
	"github.com/sugarshop/asgard-gateway/service"
	"github.com/sugarshop/asgard-gateway/util"
)

type OpenAIHandler struct {
}

func NewOpenAIHandler() *OpenAIHandler {
	return &OpenAIHandler{}
}

func (h *OpenAIHandler) Register(e *gin.Engine) {
	e.POST("/v1/openai/chat/completions", StreamWrapper(h.Completions))
}

func (h *OpenAIHandler) Completions(c *gin.Context) error {
	ctx := util.RPCContext(c)
	var reqBody model.CompletionsReqBody
	// bind json to reqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	var curCompletion model.Completion

	if !reqBody.Stream {
		response, err := service.OpenAIServiceInstance().ChatCompletion(ctx, &reqBody)
		writeResponse(c.Writer, response, &curCompletion)
		return err
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	// fixme only for test env
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	stream, apiErr := service.OpenAIServiceInstance().ChatCompletionStream(ctx, &reqBody)
	if apiErr != nil {
		return fmt.Errorf("StatusCode: %d, Type:%s, Code:%v", apiErr.HTTPStatusCode, apiErr.Type, apiErr.Code)
	}
	defer stream.Close() // defer after call for stream

	gone := c.Stream(func(w io.Writer) bool {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			//fmt.Println("[Completion]: Stream finished")
			return false
		}

		if err != nil {
			fmt.Printf("[Completion]: Stream error: %v\n", err)
			return false
		}

		return writeResponse(w, response, &curCompletion)
	})
	if gone {
		// client disconnected in middle of stream
		// do something after client is gone
		log.Println("client is gone")
	}
	reqCompletion := model.Completion{
		ChatID: curCompletion.ChatID,
		Model:  curCompletion.Model,
	}
	if len(reqBody.Messages) > 0 {
		message := reqBody.Messages[len(reqBody.Messages)-1]
		reqCompletion.Content = message.Content
		reqCompletion.Role = message.Role
	}
	if saveErr := reqCompletion.Save(); saveErr != nil {
		fmt.Printf("[Completions]: failed to save req completion: %+v", saveErr)
	}
	if saveErr := curCompletion.Save(); saveErr != nil {
		fmt.Printf("[Completions]: failed to save cur completion: %+v", saveErr)
	}

	return nil
}

func writeResponse(w io.Writer, response interface{}, completion *model.Completion) bool {
	// wrap completion.
	completion.ConstructMessage(response)

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("[writeResponse]: err:%v", err)
		return false
	}
	formattedData := fmt.Sprintf("data: %s\n\n", string(jsonBytes))
	if _, err = w.Write([]byte(formattedData)); err != nil {
		fmt.Printf("[writeResponse]: Write data:%s, error: %v", formattedData, err)
		return false
	}
	return true
}
