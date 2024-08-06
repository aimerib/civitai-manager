package actions

import (
	"civitai/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

func ModelsIndexHandler(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Initialize an empty slice to store the models
	allModels := []models.Model{}

	query := tx.Q().
		Where("id IN (SELECT model_id FROM model_versions GROUP BY model_id HAVING MAX(published_at) = published_at)").
		Order("(SELECT MAX(published_at) FROM model_versions WHERE model_versions.model_id = models.id) DESC").
		Order("id")
	err := query.Eager("ModelVersions.Images").All(&allModels)

	if err != nil {
		fmt.Println("Error getting models:", errors.Cause(err))
		return err
	}

	// Make the models available to the view
	c.Set("models", allModels)

	return c.Render(http.StatusOK, r.HTML("models/index.html"))
}
