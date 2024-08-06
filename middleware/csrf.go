package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	csrfTokenKey    = "csrf_token"
	csrfHeaderName  = "X-CSRF-Token"
	csrfContextKey  = "csrf"
	csrfTokenLength = 32
)

// CSRF returns a middleware that provides Cross-Site Request Forgery protection
func CSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" || c.Request.Method == "HEAD" || c.Request.Method == "OPTIONS" {
			ensureCSRFToken(c)
			c.Next()
			return
		}

		token := c.GetHeader(csrfHeaderName)
		if token == "" {
			token = c.PostForm(csrfTokenKey)
		}

		if token == "" || token != c.GetString(csrfTokenKey) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func ensureCSRFToken(c *gin.Context) {
	token := c.GetString(csrfTokenKey)
	if token == "" {
		token = generateCSRFToken()
		c.SetCookie(csrfTokenKey, token, 3600, "/", "", false, true)
		c.Set(csrfTokenKey, token)
	}
	c.Set(csrfContextKey, token)
}

func generateCSRFToken() string {
	b := make([]byte, csrfTokenLength)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}

// GetCSRFToken returns the CSRF token from the context
func GetCSRFToken(c *gin.Context) string {
	return c.GetString(csrfContextKey)
}
