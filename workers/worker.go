package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
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

type FetchModelsPayload struct {
	TaskID string
	Limit  int
}

func HandleFetchModelsTask(ctx context.Context, payload []byte) error {
	var p FetchModelsPayload
	if err := json.Unmarshal(payload, &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v", err)
	}

	// Simulate fetching models
	for i := 1; i <= p.Limit; i++ {
		// Check if the context has been canceled
		if ctx.Err() != nil {
			return fmt.Errorf("task canceled")
		}

		// Simulate work
		time.Sleep(time.Second)

		// Send progress update
		progress := map[string]interface{}{
			"status":       "in_progress",
			"current":      i,
			"total":        p.Limit,
			"currentModel": fmt.Sprintf("Model %d", i),
		}
		if err := SendWebSocketUpdate(p.TaskID, progress); err != nil {
			log.Printf("Failed to send WebSocket update: %v", err)
		}
	}

	// Send completion update
	complete := map[string]interface{}{
		"status": "completed",
	}
	if err := SendWebSocketUpdate(p.TaskID, complete); err != nil {
		log.Printf("Failed to send WebSocket completion update: %v", err)
	}

	return nil
}
