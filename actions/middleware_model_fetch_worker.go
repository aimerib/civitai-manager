package actions

import (
	"civitai/models"
	modelutils "civitai/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/worker"
	"github.com/gobuffalo/pop/v6"
)

// RegisterWorkers sets up our background jobs
func RegisterModelFetchWorkers(w worker.Worker, c buffalo.Context) {
	w.Register("fetch_models", func(args worker.Args) error {
		taskID := args["taskID"].(string)
		baseURL := "https://civitai.com/api/v1/models?types=Checkpoint&sort=Newest&nsfw=false&period=AllTime"

		pages := 0
		if pagesArg, ok := args["pages"]; ok && pagesArg != nil {
			pages = pagesArg.(int)
		}

		perPage := 0
		if perPageArg, ok := args["per_page"]; ok && perPageArg != nil {
			perPage = perPageArg.(int)
		}

		limit := 0
		if limitArg, ok := args["limit"]; ok && limitArg != nil {
			limit = limitArg.(int)
		}

		if limit == 0 && pages > 0 && perPage > 0 {
			limit = pages * perPage
		}

		if perPage > 0 {
			baseURL += fmt.Sprintf("&limit=%d", perPage)
		}

		totalToImport := limit
		if totalToImport == 0 {
			// If no limit is set, we can't accurately report progress
			totalToImport = 1000 // Set a default or fetch total count from API if possible
		}

		modelCount := 0
		pageCount := 0

		tx, err := pop.Connect("development")
		if err != nil {
			return err
		}

		SendWebSocketUpdate(taskID, map[string]interface{}{
			"status":       "in_progress",
			"current":      modelCount,
			"total":        totalToImport,
			"currentModel": "Fetching model list",
		})
		for url := baseURL; url != ""; {
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

			for _, model := range response.Items {
				existingModel := models.Model{}
				err := tx.Where("civitai_id = ?", model.CivitaiID).First(&existingModel)
				if err == nil {
					continue
				}

				err = modelutils.ProcessModel(tx, &model)
				if err != nil {
					return err
				}

				modelCount++
				// Report progress
				SendWebSocketUpdate(taskID, map[string]interface{}{
					"status":       "in_progress",
					"current":      modelCount,
					"total":        totalToImport,
					"currentModel": model.Name,
				})

				if limit > 0 && modelCount >= limit {
					result := StoreMessage{
						Type:    "success",
						Content: fmt.Sprintf("Background job completed successfully! Imported %d models.", modelCount),
					}
					globalMessageStore.Set(taskID, result)
					SendWebSocketUpdate(taskID, map[string]interface{}{
						"status":  "completed",
						"current": modelCount,
						"total":   totalToImport,
					})
					return nil
				}
			}

			pageCount++
			if pages > 0 && pageCount >= pages {
				result := StoreMessage{
					Type:    "success",
					Content: fmt.Sprintf("Background job completed successfully! Imported %d models.", modelCount),
				}
				globalMessageStore.Set(taskID, result)
				fmt.Printf("Reached page limit of %d pages.\n", pages)
				break
			}

			if response.Metadata.NextPage != nil {
				url = *response.Metadata.NextPage
			} else {
				url = ""
			}
		}

		result := StoreMessage{
			Type:    "success",
			Content: fmt.Sprintf("Background job completed successfully! Imported %d models.", modelCount),
		}
		globalMessageStore.Set(taskID, result)
		SendWebSocketUpdate(taskID, map[string]interface{}{
			"status":  "completed",
			"current": modelCount,
			"total":   totalToImport,
		})
		return nil
	})
}
