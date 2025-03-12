package handlers

import (
	"net/http"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
	var req struct {
		LongURL string `json:"long_url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	shortURL := services.GenerateShortURL(req.LongURL)
	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}
