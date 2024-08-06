package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SessionData makes session data available to templates
func SessionData() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		c.Set("session", session)
		c.Next()
	}
}

// GetFlash retrieves and clears a flash message from the session
func GetFlash(c *gin.Context) string {
	session := sessions.Default(c)
	flash := session.Get("flash")
	session.Delete("flash")
	session.Save()
	if flash == nil {
		return ""
	}
	return flash.(string)
}

// SetFlash sets a flash message in the session
func SetFlash(c *gin.Context, message string) {
	session := sessions.Default(c)
	session.Set("flash", message)
	session.Save()
}
