package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
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
	e.POST("/v1/openai/audio/transcriptions", JSONWrapper(h.Transcriptions))
	e.GET("/v1/openai/subscription", JSONWrapper(h.GetSubscription))

	e.POST("/v1/openai/chat/completions/api", StreamWrapper(h.CompletionsWithAPI))
}

func (h *OpenAIHandler) Completions(c *gin.Context) error {
	ctx := util.RPCContext(c)
	var reqBody model.CompletionsReqBody
	// bind json to reqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	tokenSufficient, err := service.ChattyAIServiceInstance().TokenIsSufficient(ctx, reqBody.UID)
	if err != nil {
		log.Println("[Completions]: TokenIsSufficient err ", err)
		return err
	}
	if !tokenSufficient {
		return fmt.Errorf("token insufficient")
	}
	return h.completions(c, &reqBody)
}

// CompletionsWithAPI Completions from api request
func (h *OpenAIHandler) CompletionsWithAPI(c *gin.Context) error {
	ctx := util.RPCContext(c)
	var reqBody model.CompletionsReqBody
	// bind json to reqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	api_access, err := service.ChattyAIServiceInstance().APIAccess(ctx, reqBody.APIKey)
	if err != nil {
		log.Println("[Completions]: APIAccess err ", err)
		return err
	}
	if !api_access {
		return fmt.Errorf("access denied, please check quota")
	}
	// using api_key as uid
	reqBody.UID = reqBody.APIKey
	return h.completions(c, &reqBody)
}

func (h *OpenAIHandler) completions(c *gin.Context, reqBody *model.CompletionsReqBody) error {
	ctx := util.RPCContext(c)

	var curCompletion model.Completion

	if !reqBody.Stream {
		response, err := service.OpenAIServiceInstance().ChatCompletion(ctx, reqBody)
		writeResponse(c.Writer, response, &curCompletion)
		return err
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	// fixme only for test env
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	stream, apiErr := service.OpenAIServiceInstance().ChatCompletionStream(ctx, reqBody)
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

	// update tokens
	service.OpenAIServiceInstance().CountStreamTokens(reqBody, curCompletion.Content)
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

// Transcriptions transcription from audio to text.
func (h *OpenAIHandler) Transcriptions(c *gin.Context) (interface{}, error) {
	ctx := util.RPCContext(c)
	// bind json to reqBody
	model, err := util.String(c, "model")
	if err != nil {
		log.Println("[Transcriptions]: get model from param err", err)
		return nil, err
	}
	prompt, err := util.String(c, "prompt")
	if err != nil {
		log.Println("[Transcriptions]: get prompt from param err", err)
		return nil, err
	}
	temperature, err := util.Float32(c, "temperature")
	if err != nil {
		log.Println("[Transcriptions]: get temperature from param err", err)
		return nil, err
	}
	language, err := util.String(c, "language")
	if err != nil {
		log.Println("[Transcriptions]: get language from param err", err)
		return nil, err
	}
	_, headers, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get file failed"})
		return nil, err
	}

	// save file using filename
	filepath := "./" + headers.Filename
	err = c.SaveUploadedFile(headers, filepath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "save file failed"})
		return nil, err
	}
	defer os.Remove(filepath)

	return service.OpenAIServiceInstance().Transcriptions(ctx, &openai.AudioRequest{
		Model:       model,
		FilePath:    filepath,
		Prompt:      prompt,
		Temperature: temperature,
		Language:    language,
		Format:      "json",
	})
}

// GetSubscription get ChattyAI Quota
func (h *OpenAIHandler) GetSubscription(c *gin.Context) (interface{}, error) {
	ctx := util.RPCContext(c)
	uid, err := util.String(c, "uid")
	if err != nil {
		log.Println("[GetSubscription]: get uid from param err", err)
		return nil, err
	}
	rights, err := service.ChattyAIServiceInstance().GetSubscriptionByUID(ctx, uid)
	if err != nil {
		log.Println("[GetSubscription]: GetSubscriptionByUID err", err)
		return nil, err
	}
	return rights, nil
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
