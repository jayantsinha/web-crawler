package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	var response map[string]string

	handler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}

	router := gin.New()
	router.GET("/ping", handler)

	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	err := json.Unmarshal([]byte(resp.Body.String()), &response)
	value, exists := response["message"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "pong", value)

}
