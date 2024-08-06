package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SettingsHandler struct{}

func NewSettingsHandler() *SettingsHandler {
	return &SettingsHandler{}
}

func (h *SettingsHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "settings/index.html", gin.H{
		"title": "Settings",
	})
}
