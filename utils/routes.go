package utils

import (
	"civitai-manager/views"
	"fmt"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func RenderView(c *gin.Context, view templ.Component) {
	fmt.Println("Rendering view:", view)
	c.Set("content", view)
}

func WithLayout(t string, h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c)
		if !c.IsAborted() {
			content := c.MustGet("content").(templ.Component)
			csrfField := csrf.TemplateField(c.Request)
			templWrapper := fmt.Sprint("%v", csrfField)
			views.Layout(t, content, templWrapper).Render(c, c.Writer)
		}
	}
}
