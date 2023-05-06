package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/model"
	"io"
	"net/http"
)

type OpenAIHandler struct {
}

func NewOpenAIHandler() *OpenAIHandler {
	return &OpenAIHandler{}
}

func (h *OpenAIHandler) Register(e *gin.Engine)  {
	e.POST("/v1/openai/chat/completions", h.Completions)
}

func (h *OpenAIHandler) Completions(c *gin.Context) {
	var reqBody model.CompletionsReqBody
	// bind json to reqBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.OpenAIStream(c, reqBody)
	if err != nil {
		var openaiErr *model.OpenAIError
		if errors.As(err, &openaiErr) {
			c.JSON(http.StatusBadRequest, gin.H{"error": openaiErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	defer resp.Body.Close()

	// todo: fix stream logic below.
	c.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, resp.Body)
		if err != nil {
			return false
		}
		return false
	})
}


func (h *OpenAIHandler) OpenAIStream(c *gin.Context, body model.CompletionsReqBody) (*http.Response, error) {
	m := body.Model
	systemPrompt := body.SystemPrompt
	temperature := body.Temperature
	key := body.Key
	messages := body.Messages
	url := fmt.Sprintf("%s/v1/chat/completions", model.OPENAIAPIHOST)
	if model.OPENAIAPITYPE == "azure" {
		url = fmt.Sprintf("%s/openai/deployments/%s/chat/completions?api-version=%s", model.OPENAIAPIHOST, model.AZUREDEPLOYMENTID, model.OPENAIAPIVERSION)
	}

	reqBody := model.ChatRequest{
		Model:       m.ID,
		Messages:    append([]model.Message{{Role: "system", Content: systemPrompt}}, messages...),
		MaxTokens:   1000,
		Temperature: temperature,
		Stream:      true,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if model.OPENAIAPITYPE == "openai" {
		if key != "" {
			req.Header.Set("Authorization", "Bearer "+key)
		} else {
			req.Header.Set("Authorization", "Bearer "+model.OPENAIAPIKEY)
		}
	}
	if model.OPENAIAPITYPE == "azure" {
		if key != "" {
			req.Header.Set("api-key", key)
		} else {
			req.Header.Set("api-key", model.OPENAIAPIKEY)
		}
	}
	if model.OPENAIAPITYPE == "openai" && model.OPENAIORGANIZATION != "" {
		req.Header.Set("OpenAI-Organization", model.OPENAIORGANIZATION)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)
		var result struct {
			Error *model.OpenAIError `json:"error"`
		}
		if err := json.Unmarshal(bodyBytes, &result); err != nil {
			return nil, fmt.Errorf("OpenAI API returned an error: %s", resp.Status)
		}
		if result.Error != nil {
			return nil, result.Error
		}
		return nil, fmt.Errorf("OpenAI API returned an error: %s", resp.Status)
	}

	return resp, nil
}
