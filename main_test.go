package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPing(t *testing.T) {
	resp := gin.H{"message": "pong"}
	router := setupRouter()

	w := performRequest(router, http.MethodGet, "/ping")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, resp["message"], value)

}
