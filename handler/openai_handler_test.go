package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/sugarshop/asgard-gateway/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenAIHandler_Completions(t *testing.T) {
	// Create a new instance of the OpenAIHandler struct
	h := NewOpenAIHandler()

	// Create a new HTTP recorder (to record the response)
	rec := CreateTestResponseRecorder()

	// create a new gin engine instance
	_, router := gin.CreateTestContext(rec)
	// register the h with the router
	h.Register(router)

	// Create a new HTTP request
	reqBody := model.CompletionsReqBody{
		Model: model.OpenAIModel{ID: "gpt-3.5-turbo"},
		SystemPrompt: model.NEXTPUBLICDEFAULTSYSTEMPROMPT,
		Temperature: 1,
		Key: model.OPENAIAPIKEY,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "who are u?"},
		},
	}
	reqBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/v1/openai/chat/completions", bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}


	// perform the test request
	router.ServeHTTP(rec, req)

	// check the response status code
	if rec.Code != http.StatusOK {
		t.Fatalf("Unexpected response status code: %d", rec.Code)
	}

	// check the response body
	var buf bytes.Buffer
	_, err = io.Copy(&buf, rec.Body)
	fmt.Println(buf.String())
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if buf.Len() == 0 {
		t.Fatalf("Empty response body")
	}
}

type TestResponseRecorder struct {
	*httptest.ResponseRecorder
	closeChannel chan bool
}

func (r *TestResponseRecorder) CloseNotify() <-chan bool {
	return r.closeChannel
}

func CreateTestResponseRecorder() *TestResponseRecorder {
	return &TestResponseRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}