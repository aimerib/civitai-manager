package handlers

import (
	"civitai-manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ModelHandler struct {
	DB *gorm.DB
}

func NewModelHandler(db *gorm.DB) *ModelHandler {
	return &ModelHandler{DB: db}
}

func (h *ModelHandler) ModelsIndex(c *gin.Context) {
	var allModels []models.Model

	err := h.DB.
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

	c.HTML(http.StatusOK, "models/index.html", gin.H{
		"models": allModels,
	})
}

func (h *ModelHandler) ModelsShow(c *gin.Context) {
	id := c.Param("id")

	var model models.Model
	err := h.DB.
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

	c.HTML(http.StatusOK, "models/show.html", gin.H{
		"model": model,
	})
}
