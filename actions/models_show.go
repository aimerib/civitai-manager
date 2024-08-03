package actions

import (
	"civitai/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func ModelsShowHandler(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	id := c.Param("id")

	// Initialize an empty slice to store the models
	model := models.Model{}

	// Get the model with the given ID from the DB
	err := tx.Where("id =?", id).First(&model)
	if err != nil {
		return c.Error(404, fmt.Errorf("models.Model with ID=%s not found", id))
	}

	// Make the models available to the view
	c.Set("model", model)

	return c.Render(http.StatusOK, r.HTML("models/show.html"))
}
