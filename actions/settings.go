package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func SettingsHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("settings/index.plush.html"))
}
