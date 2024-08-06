package helpers

import (
	"fmt"
	"html/template"
	"path/filepath"
)

type TemplateHelpers struct{}

func NewTemplateHelpers() *TemplateHelpers {
	return &TemplateHelpers{}
}

func (th *TemplateHelpers) FuncMap() template.FuncMap {
	return template.FuncMap{
		"assetPath": th.AssetPath,
		"cssPath":   th.CSSPath,
		"jsPath":    th.JSPath,
		"imagePath": th.ImagePath,
	}
}

func (th *TemplateHelpers) AssetPath(name string) string {
	return fmt.Sprintf("/public/%s", name)
}

func (th *TemplateHelpers) CSSPath(name string) string {
	return th.AssetPath(filepath.Join("css", name))
}

func (th *TemplateHelpers) JSPath(name string) string {
	return th.AssetPath(filepath.Join("js", name))
}

func (th *TemplateHelpers) ImagePath(name string) string {
	return th.AssetPath(filepath.Join("images", name))
}
