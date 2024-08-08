package workers

import (
	"civitai-manager/models"
	"civitai-manager/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Type    string
	Payload []byte
}

type Worker struct {
	tasks    chan Task
	quit     chan bool
	handlers map[string]func(context.Context, []byte) error
	wg       sync.WaitGroup
}

func NewWorker(numWorkers int) *Worker {
	return &Worker{
		tasks:    make(chan Task, 100),
		quit:     make(chan bool),
		handlers: make(map[string]func(context.Context, []byte) error),
	}
}

func (w *Worker) Start() {
	for i := 0; i < cap(w.tasks); i++ {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for {
				select {
				case task := <-w.tasks:
					if handler, ok := w.handlers[task.Type]; ok {
						if err := handler(context.Background(), task.Payload); err != nil {
							log.Printf("Error processing task: %v", err)
						}
					} else {
						log.Printf("No handler for task type: %s", task.Type)
					}
				case <-w.quit:
					return
				}
			}
		}()
	}
}

func (w *Worker) Stop() {
	close(w.quit)
	w.wg.Wait()
}

func (w *Worker) Enqueue(taskType string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.tasks <- Task{Type: taskType, Payload: data}
	return nil
}

func (w *Worker) RegisterHandler(taskType string, handler func(context.Context, []byte) error) {
	w.handlers[taskType] = handler
}

const TypeFetchModels = "fetch:models"

// type FetchModelsPayload struct {
// 	TaskID string
// 	Limit  int
// }

// func HandleFetchModelsTask(ctx context.Context, payload []byte) error {
// 	var p FetchModelsPayload
// 	if err := json.Unmarshal(payload, &p); err != nil {
// 		return fmt.Errorf("json.Unmarshal failed: %v", err)
// 	}

// 	// Simulate fetching models
// 	for i := 1; i <= p.Limit; i++ {
// 		// Check if the context has been canceled
// 		if ctx.Err() != nil {
// 			return fmt.Errorf("task canceled")
// 		}

// 		// Simulate work
// 		time.Sleep(time.Second)

// 		// Send progress update
// 		progress := map[string]interface{}{
// 			"status":       "in_progress",
// 			"current":      i,
// 			"total":        p.Limit,
// 			"currentModel": fmt.Sprintf("Model %d", i),
// 		}
// 		if err := handlers.SendWebSocketUpdate(p.TaskID, progress); err != nil {
// 			log.Printf("Failed to send WebSocket update: %v", err)
// 		}
// 	}

// 	// Send completion update
// 	complete := map[string]interface{}{
// 		"status": "completed",
// 	}
// 	if err := handlers.SendWebSocketUpdate(p.TaskID, complete); err != nil {
// 		log.Printf("Failed to send WebSocket completion update: %v", err)
// 	}

// 	return nil
// }

type FetchModelsPayload struct {
	TaskID  string
	Pages   int
	PerPage int
	Limit   int
}

func (w *Worker) RegisterModelFetchWorker(db *gorm.DB) {
	modelProcessor := services.NewModelProcessor(db)

	w.RegisterHandler(TypeFetchModels, func(ctx context.Context, payload []byte) error {
		var args FetchModelsPayload
		if err := json.Unmarshal(payload, &args); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v", err)
		}

		pages := 0
		if args.Pages > 0 {
			pages = args.Pages
		}

		perPage := 0
		if args.PerPage > 0 {
			perPage = args.PerPage
		}

		limit := 0
		if args.Limit > 0 {
			limit = args.Limit
		}

		baseURL := "https://civitai.com/api/v1/models?types=Checkpoint&sort=Newest&nsfw=false&period=AllTime"

		if perPage > 0 {
			baseURL += fmt.Sprintf("&limit=%d", perPage)
		}

		totalToImport := limit
		if totalToImport == 0 && pages > 0 && perPage > 0 {
			totalToImport = pages * perPage
		}
		if totalToImport == 0 {
			totalToImport = 1000 // Default value
		}

		modelCount := 0
		pageCount := 0

		SendWebSocketUpdate(args.TaskID, map[string]interface{}{
			"status":       "in_progress",
			"current":      modelCount,
			"total":        totalToImport,
			"currentModel": "Fetching model list",
		})

		for url := baseURL; url != ""; {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
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
					var existingModel models.Model
					err := db.Where("civitai_id = ?", model.CivitaiID).First(&existingModel).Error
					if err == nil {
						continue
					}

					err = modelProcessor.ProcessModel(&model)
					if err != nil {
						return err
					}

					modelCount++
					SendWebSocketUpdate(args.TaskID, map[string]interface{}{
						"status":       "in_progress",
						"current":      modelCount,
						"total":        totalToImport,
						"currentModel": model.Name,
					})

					if limit > 0 && modelCount >= limit {
						SendWebSocketUpdate(args.TaskID, map[string]interface{}{
							"status":  "completed",
							"current": modelCount,
							"total":   totalToImport,
						})
						return nil
					}
				}

				pageCount++
				if pages > 0 && pageCount >= pages {
					break
				}

				if response.Metadata.NextPage != nil {
					url = *response.Metadata.NextPage
				} else {
					url = ""
				}

				// Add a small delay to avoid overwhelming the API
				time.Sleep(time.Second)
			}
		}

		SendWebSocketUpdate(args.TaskID, map[string]interface{}{
			"status":  "completed",
			"current": modelCount,
			"total":   totalToImport,
		})
		return nil
	})
}

func SendWebSocketUpdate(taskID string, message interface{}) {
	// Implementation depends on your WebSocket setup
	// This is a placeholder - replace with your actual implementation
	log.Printf("WebSocket Update for task %s: %v", taskID, message)
}
