package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenAIHandler_Completions(t *testing.T) {
	handler := NewOpenAIHandler()

	// Create a test Gin context
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/openai/chat/completions", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the handler function
	result, err := handler.Completions(c)

	// Check the result and ensure no error occurred
	assert.Nil(t, err, "Unexpected error")
	assert.Nil(t, result, "Unexpected response body")
}
