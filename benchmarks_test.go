package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)
type mockWriter struct {
	headers http.Header
}

func (m *mockWriter) Header() (h http.Header) {
	return m.headers
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockWriter) WriteHeader(int) {}

func newMockWriter() *mockWriter {
	return &mockWriter{
		http.Header{},
	}
}

func BenchmarkScrapingController(b *testing.B) {
	router := gin.New()
	router.GET("/crawl", ScrapingController)
	req, _ := http.NewRequest(http.MethodGet, "/crawl", nil)
	req.Header.Add("Scrape", "http://testing-ground.scraping.pro/blocks")
	w := newMockWriter()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(w, req)
	}
}