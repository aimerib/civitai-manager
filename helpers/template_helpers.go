package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/csrf"
)

type TemplateHelpers struct{}

func NewTemplateHelpers() *TemplateHelpers {
	return &TemplateHelpers{}
}

func (th *TemplateHelpers) FuncMap() template.FuncMap {
	return template.FuncMap{
		"assetPath":  th.AssetPath,
		"cssPath":    th.CSSPath,
		"jsPath":     th.JSPath,
		"imagePath":  th.ImagePath,
		"safeHTML":   th.SafeHTML,
		"csrfField":  th.CsrfField,
		"flashIcon":  th.FlashIcon,
		"flashTitle": th.FlashTitle,
		"contains":   strings.Contains,
		"csrfToken":  th.CsrfToken,
	}
}
func (th *TemplateHelpers) CsrfField(r *http.Request) template.HTML {
	return csrf.TemplateField(r)
}

func (th *TemplateHelpers) CsrfToken(r *http.Request) string {
	return csrf.Token(r)
}

func (th *TemplateHelpers) AssetPath(name string) string {
	return fmt.Sprintf("/public/assets/%s", name)
}

func (th *TemplateHelpers) CSSPath(name string) template.HTML {
	return template.HTML(`<link rel="stylesheet" href="` + th.AssetPath(filepath.Join("css", name)) + `" />`)
}

func (th *TemplateHelpers) JSPath(name string) template.HTML {
	return template.HTML(`<script src="` + th.AssetPath(filepath.Join("js", name)) + `"></script>`)
}

func (th *TemplateHelpers) ImagePath(name string) template.HTML {
	return template.HTML(`<img src="` + th.AssetPath(filepath.Join("images", name)) + `" alt="` + name + `" />`)
}

func (th *TemplateHelpers) SafeHTML(s string) template.HTML {
	return template.HTML(s)
}

func (th *TemplateHelpers) FlashIcon(flashType string) template.HTML {
	var icon string
	switch flashType {
	case "success":
		icon = `<svg class="h-6 w-6 text-green-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
				  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
				</svg>`
	case "error":
		icon = `<svg class="h-6 w-6 text-red-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
				  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
				</svg>`
	case "warning":
		icon = `<svg class="h-6 w-6 text-yellow-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
				  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
				</svg>`
	default:
		icon = `<svg class="h-6 w-6 text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
				  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
				</svg>`
	}
	return template.HTML(icon)
}

func (th *TemplateHelpers) FlashTitle(flashType string) string {
	switch flashType {
	case "success":
		return "Success"
	case "error":
		return "Error"
	case "warning":
		return "Warning"
	default:
		return "Information"
	}
}
