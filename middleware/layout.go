package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// SetLayout middleware sets the appropriate layout based on the request
func SetLayout() gin.HandlerFunc {
	return func(c *gin.Context) {
		layout := "layouts/application.html"
		if c.GetHeader("HX-Request") == "true" && strings.Trim(c.Request.URL.Path, "/") != "" {
			layout = "layouts/content.html"
		}
		c.Set("layout", layout)
		c.Next()
	}
}
