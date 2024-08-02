package grifts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"civitai/models"

	"github.com/gobuffalo/grift/grift"
	"github.com/gobuffalo/pop/v6"
)

var _ = grift.Namespace("api", func() {
	grift.Desc("fetch_models", "Fetch and update models from API")
	grift.Add("fetch_models", func(c *grift.Context) error {
		baseURL := "https://civitai.com/api/v1/models?limit=10&types=Checkpoint"

		// Get database connection
		tx, err := pop.Connect("development")
		if err != nil {
			return err
		}

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
					// Model already exists, stop processing
					return nil
				}

				err = processModel(tx, &model)
				if err != nil {
					return err
				}
			}
			fmt.Println(response)

			// Set next page URL
			if response.Metadata.NextPage != nil {
				url = *response.Metadata.NextPage
			} else {
				url = ""
			}
		}

		return nil
	})
})

func processModel(tx *pop.Connection, model *models.Model) error {
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
	model.Stat.ModelID = model.ID
	// Process model stats
	return tx.Create(&model.Stat)
}

func processCreator(tx *pop.Connection, model *models.Model) error {
	model.Creator.ModelID = model.ID
	return tx.Create(&model.Creator)
}

func processFiles(tx *pop.Connection, modelVersion *models.ModelVersion) error {
	for _, file := range modelVersion.Files {
		file.ModelVersionID = modelVersion.ID
		err := tx.Create(&file)
		if err != nil {
			return err
		}
	}
	return nil
}

func processImages(tx *pop.Connection, modelVersion *models.ModelVersion) error {
	for _, image := range modelVersion.Images {
		image.ModelVersionID = modelVersion.ID
		err := tx.Create(&image)
		if err != nil {
			return err
		}
	}
	return nil
}

func processModelVersionStats(tx *pop.Connection, modelVersion *models.ModelVersion) error {
	modelVersion.Stats.ModelVersionID = modelVersion.ID
	return tx.Create(&modelVersion.Stats)
}
