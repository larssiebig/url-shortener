package test

import (
	"testing"

	"github.com/larssiebig/url-shortener/internal/services"
)

func TestShortenURL(t *testing.T) {
	shortURL := services.GenerateShortCode("https://golang.org")
	if len(shortURL) != 6 {
		t.Errorf("Expected short URL length of 6, got %d", len(shortURL))
	}
}
