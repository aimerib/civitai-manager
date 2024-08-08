package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// CacheMiddleware sets caching headers for image requests
func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/images/") {
			c.Header("Cache-Control", "public, max-age=31536000")
		}
		c.Next()
	}
}
