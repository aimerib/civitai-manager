package handlers

import (
	"civitai-manager/utils"
	"civitai-manager/views"
	"civitai-manager/workers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SettingsHandler struct {
	Worker *workers.Worker
}

func NewSettingsHandler(worker *workers.Worker) *SettingsHandler {
	return &SettingsHandler{Worker: worker}
}

func (h *SettingsHandler) Index(c *gin.Context) {
	utils.RenderView(c, views.SettingsIndex())
	// c.HTML(http.StatusOK, "settings/index", gin.H{"Request": c.Request})
}

func (h *SettingsHandler) StartBackgroundFetchJob(c *gin.Context) {
	taskID := uuid.New().String()
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	perPage, _ := strconv.Atoi(c.PostForm("per_page"))
	pages, _ := strconv.Atoi(c.PostForm("pages"))

	payload := workers.FetchModelsPayload{
		TaskID:  taskID,
		Limit:   limit,
		PerPage: perPage,
		Pages:   pages,
	}

	err := h.Worker.Enqueue(workers.TypeFetchModels, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"taskID": taskID, "Request": c.Request})
}
