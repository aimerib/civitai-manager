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
	// existingModel := models.Model{}
	// err := tx.Where("civitai_id = ?", model.CivitaiID).First(&existingModel)

	// if err == nil {
	// 	// Update existing model
	// 	existingModel = *model
	// 	err = tx.Update(&existingModel)
	// } else {
	// 	// Create new model
	// 	err = tx.Create(model)
	// }

	// if err != nil {
	// 	return err
	// }
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

	// err = processTrainedWords(tx, model)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func processModelVersions(tx *pop.Connection, model *models.Model) error {
	for _, version := range model.ModelVersions {
		version.ModelID = model.ID
		err := tx.Create(&version)
		if err != nil {
			return err
		}

		// Process files, stats, images, hashes
		// Add similar processing for these associated data
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

// func processTrainedWords(tx *pop.Connection, model *models.Model) error {
// 	// Process trained words
// 	for _, word := range model. {
// 		trainedWord := models.TrainedWord{
// 			ModelID: model.ID,
// 			Word:    word,
// 		}
// 		err := tx.Create(&trainedWord)
// 		if err != nil {
// 			return err
// 		}
// 	}
// return nil
// }
