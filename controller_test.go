package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestScrapingController(t *testing.T) {

	router := setupRouter()
	w := performRequest(router, "GET", "/crawl")
	assert.Equal(t, http.StatusOK, w.Code)

	var response JsonResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	if err != nil {
		t.Errorf("Invalid response format: %v", err)
	}

	if response.Urls[0].Loc == "" {
		t.Errorf("Empty response from /crawl endpoint!")
	}

}
