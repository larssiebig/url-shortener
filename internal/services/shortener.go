package services

import (
	"crypto/sha256"
	"encoding/base64"
)

// GenerateShortCode creates a short identifier for a URL
func GenerateShortCode(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))
	shortCode := base64.URLEncoding.EncodeToString(hash[:])[:6]
	return shortCode
}
