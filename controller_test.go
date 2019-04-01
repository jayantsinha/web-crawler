package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrapingController(t *testing.T) {
	response := make(map[string]interface{})
	router := SetupRouter()
	router.GET("/crawl", ScrapingController)

	req, _ := http.NewRequest(http.MethodGet, "/crawl", nil)
	req.Header.Add("Scrape", "http://testing-ground.scraping.pro/blocks")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	err := json.Unmarshal([]byte(resp.Body.String()), &response)
	respLen := len(response)
	assert.Nil(t, err)
	if respLen == 0 {

		t.Errorf("Expecting response length to be > 0 but got %v", respLen)
	}
}
