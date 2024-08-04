package grifts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"civitai/models"

	"github.com/gobuffalo/grift/grift"
	"github.com/gobuffalo/pop/v6"
)

var _ = grift.Namespace("api", func() {
	grift.Desc("fetch_models", "Fetch and update models from API. Usage: buffalo task api:fetch_models [limit]")
	grift.Add("fetch_models", func(c *grift.Context) error {
		// baseURL := "https://civitai.com/api/v1/models?limit=100&types=Checkpoint&tags=nsfw&sort=Newest&nsfw=true&period=AllTime"
		baseURL := "https://civitai.com/api/v1/models?limit=100&types=Checkpoint&sort=Newest&nsfw=false&period=AllTime"

		// Parse the optional limit argument
		var limit int
		if len(c.Args) > 0 {
			_, err := fmt.Sscanf(c.Args[0], "%d", &limit)
			if err != nil {
				return fmt.Errorf("invalid limit argument: please provide a number")
			}
		}

		// Get database connection
		tx, err := pop.Connect("development")
		if err != nil {
			return err
		}

		modelCount := 0

		for url := baseURL; url != ""; {
			// Fetch JSON data
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			var response struct {
				Items    []models.Model `json:"items"`
				Metadata struct {
					NextPage *string `json:"nextPage"`
				} `json:"metadata"`
			}

			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				return err
			}

			// Process each model
			for _, model := range response.Items {
				existingModel := models.Model{}
				err := tx.Where("civitai_id = ?", model.CivitaiID).First(&existingModel)
				if err == nil {
					// Model already exists, skip processing
					return nil
				}

				err = processModel(tx, &model)
				if err != nil {
					return err
				}

				modelCount++
				if model.Name != nil {
					fmt.Printf("Imported model: %s (ID: %d)\n", *model.Name, model.CivitaiID)
				} else {
					fmt.Printf("Imported model with no name (ID: %d)\n", model.CivitaiID)
				}
				// Check if we've reached the limit
				if limit > 0 && modelCount >= limit {
					fmt.Printf("Reached import limit of %d models.\n", limit)
					return nil
				}
			}

			// Set next page URL
			if response.Metadata.NextPage != nil {
				url = *response.Metadata.NextPage
			} else {
				url = ""
			}
		}

		fmt.Printf("Imported a total of %d models.\n", modelCount)
		return nil
	})
})

func processModel(tx *pop.Connection, model *models.Model) error {
	if model.Description != nil {
		re := regexp.MustCompile(`\s*id="[^"]*"`)
		*model.Description = re.ReplaceAllString(*model.Description, "")
	}

	err := tx.Create(model)
	if err != nil {
		return err
	}

	// Process associated data
	err = processModelVersions(tx, model)
	if err != nil {
		return err
	}

	err = processTags(tx, model)
	if err != nil {
		return err
	}

	err = processStats(tx, model)
	if err != nil {
		return err
	}

	err = processCreator(tx, model)
	if err != nil {
		return err
	}

	return nil
}

func processModelVersions(tx *pop.Connection, model *models.Model) error {
	for _, version := range model.ModelVersions {
		version.ModelID = model.ID
		err := tx.Create(&version)
		if err != nil {
			return err
		}

		err = processFiles(tx, &version)
		if err != nil {
			return err
		}

		err = processImages(tx, &version)
		if err != nil {
			return err
		}

		err = processModelVersionStats(tx, &version)
		if err != nil {
			return err
		}
	}
	return nil
}

func processTags(tx *pop.Connection, model *models.Model) error {
	for _, tag := range model.Tags {
		err := tag.FindOrCreate(tx)
		if err != nil {
			return err
		}

		modelTag := models.ModelTag{
			ModelID: model.ID,
			TagID:   tag.ID,
		}
		err = tx.Create(&modelTag)
		if err != nil {
			return err
		}
	}
	return nil
}

func processStats(tx *pop.Connection, model *models.Model) error {
	model.Stats.ModelID = model.ID
	// Process model stats
	return tx.Create(&model.Stats)
}

func processCreator(tx *pop.Connection, model *models.Model) error {
	model.Creator.ModelID = model.ID
	return tx.Create(&model.Creator)
}

func processFiles(tx *pop.Connection, modelVersion *models.ModelVersions) error {
	for _, file := range modelVersion.Files {
		file.ModelVersionsID = modelVersion.ID
		err := tx.Create(&file)
		if err != nil {
			return err
		}
	}
	return nil
}

func processImages(tx *pop.Connection, modelVersion *models.ModelVersions) error {
	for _, image := range modelVersion.Images {
		image.ModelVersionsID = modelVersion.ID
		err := tx.Create(&image)
		if err != nil {
			return err
		}
	}
	return nil
}

func processModelVersionStats(tx *pop.Connection, modelVersion *models.ModelVersions) error {
	modelVersion.Stats.ModelVersionsID = modelVersion.ID
	return tx.Create(&modelVersion.Stats)
}
