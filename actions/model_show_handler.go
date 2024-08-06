package actions

import (
	"civitai/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func ModelsShowHandler(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	id := c.Param("id")

	model := models.Model{}

	err := tx.Find(&model, id)
	if err != nil {
		c.Logger().Errorf("Error fetching base model: %v", err)
		return c.Error(404, err)
	}
	err = tx.Load(&model, "Creator")
	if err != nil {
		c.Logger().Errorf("Error loading Creator: %v", err)
	}
	if err != nil {
		c.Logger().Errorf("Error loading Stats: %v", err)
	}
	err = tx.Load(&model, "ModelVersions")
	if err != nil {
		c.Logger().Errorf("Error loading ModelVersions: %v", err)
	}

	if len(model.ModelVersions) > 0 {
		for i := range model.ModelVersions {
			err = tx.Load(&model.ModelVersions[i], "Images")
			if err != nil {
				c.Logger().Errorf("Error loading Images for ModelVersion %d: %v", i, err)
			}

			err = tx.Load(&model.ModelVersions[i], "Files")
			if err != nil {
				c.Logger().Errorf("Error loading Files for ModelVersion %d: %v", i, err)
			}
		}
	}

	c.Logger().Infof("Retrieved model with civitai-id: %+v", model.CivitaiID)
	c.Set("model", model)

	return c.Render(http.StatusOK, r.HTML("models/show.html"))
}
