package handlers

import (
	"civitai-manager/middleware"
	"civitai-manager/models"
	"civitai-manager/utils"
	"civitai-manager/views"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ModelHandler struct{}

func NewModelHandler() *ModelHandler {
	return &ModelHandler{}
}

func (h *ModelHandler) ModelsIndex(c *gin.Context) {
	var allModels []models.Model
	db := middleware.GetTrx(c)
	err := db.
		Preload("ModelVersions.Images").
		Joins("JOIN model_versions ON models.id = model_versions.model_id").
		Group("models.id").
		Order("MAX(model_versions.published_at) DESC").
		Order("models.id").
		Find(&allModels).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting models"})
		return
	}

	// c.HTML(http.StatusOK, "models/index", gin.H{
	// 	"models":  allModels,
	// 	"Request": c.Request,
	// })
	c.Header("Content-Type", "text/html")
	utils.RenderView(c, views.ModelsIndex(allModels))
	// views.Layout("Models", views.ModelsIndex(allModels)).Render(c, c.Writer)
}

func (h *ModelHandler) ModelsShow(c *gin.Context) {
	id := c.Param("id")
	db := middleware.GetTrx(c)
	var model models.Model
	err := db.
		Preload("Creator").
		Preload("Stats").
		Preload("ModelVersions.Images").
		Preload("ModelVersions.Files").
		First(&model, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching model"})
		}
		return
	}

	utils.RenderView(c, views.ModelsShow(model))

	// c.HTML(http.StatusOK, "models/show", gin.H{
	// 	"model":   model,
	// 	"Request": c.Request,
	// })
}
