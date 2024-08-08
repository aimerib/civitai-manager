package main

import (
	"civitai-manager/config"
	"civitai-manager/models"
	"civitai-manager/services"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	pages := flag.Int("pages", 0, "Number of pages to fetch")
	perPage := flag.Int("per_page", 0, "Number of items per page")
	limit := flag.Int("limit", 0, "Total number of models to fetch")
	flag.Parse()

	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	modelProcessor := services.NewModelProcessor(db)

	baseURL := "https://civitai.com/api/v1/models?types=Checkpoint&sort=Newest&nsfw=false&period=AllTime"

	if *limit == 0 && *pages > 0 && *perPage > 0 {
		*limit = *pages * *perPage
	}

	if *perPage > 0 {
		baseURL += fmt.Sprintf("&limit=%d", *perPage)
	}

	modelCount := 0
	pageCount := 0

	for url := baseURL; url != ""; {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to fetch URL: %v\n", err)
			return
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
			fmt.Printf("Failed to decode response: %v\n", err)
			return
		}

		for _, model := range response.Items {
			var existingModel models.Model
			err := db.Where("civitai_id = ?", model.CivitaiID).First(&existingModel).Error
			if err == nil {
				continue
			}

			err = modelProcessor.ProcessModel(&model)
			if err != nil {
				fmt.Printf("Failed to process model: %v\n", err)
				return
			}

			modelCount++
			if model.Name != "" {
				fmt.Printf("Imported model: %s (ID: %d)\n", model.Name, model.CivitaiID)
			} else {
				fmt.Printf("Imported model with no name (ID: %d)\n", model.CivitaiID)
			}

			if *limit > 0 && modelCount >= *limit {
				fmt.Printf("Reached import limit of %d models.\n", *limit)
				return
			}
		}

		pageCount++
		if *pages > 0 && pageCount >= *pages {
			fmt.Printf("Reached page limit of %d pages.\n", *pages)
			break
		}

		if response.Metadata.NextPage != nil {
			url = *response.Metadata.NextPage
		} else {
			url = ""
		}
	}

	fmt.Printf("Imported a total of %d models across %d pages.\n", modelCount, pageCount)
}
