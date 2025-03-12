package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/larssiebig/url-shortener/internal/repository"
	"github.com/larssiebig/url-shortener/internal/services"
)

// ShortenURL handles URL shortening requests
func ShortenURL(c *gin.Context) {
	var req struct {
		LongURL string `json:"long_url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortCode := services.GenerateShortCode(req.LongURL)
	repository.SaveURL(shortCode, req.LongURL)

	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortCode})
}

// RedirectURL handles redirection from short URL to original URL
func RedirectURL(c *gin.Context) {
	shortCode := c.Param("shortcode")
	longURL := repository.GetOriginalURL(shortCode)

	if longURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	c.Redirect(http.StatusFound, longURL)
}
