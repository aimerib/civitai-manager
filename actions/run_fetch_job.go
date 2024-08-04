package actions

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/grift/grift"
)

func RunFetchJobHandler(c buffalo.Context) error {
	// Set up SSE
	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")

	flusher, ok := c.Response().(http.Flusher)
	if !ok {
		return c.Error(500, errors.New("streaming not supported"))
	}

	resultChan := make(chan string)

	// Start the background task
	go func() {
		context := grift.NewContext("fetch_models")
		context.Set("args", []string{"--limit=100"})
		err := grift.Run("api:fetch_models", context)
		if err != nil {
			c.Logger().Error(err)
			resultChan <- "error"
		} else {
			resultChan <- "success"
		}
	}()

	// Send initial processing status
	fmt.Fprintf(c.Response(), "data: processing\n\n")
	flusher.Flush()

	// Wait for result or timeout
	select {
	case result := <-resultChan:
		var status string
		switch result {
		case "success":
			status = "completed"
		case "error":
			status = "failed"
		}
		fmt.Fprintf(c.Response(), "data: %s\n\n", status)
	case <-time.After(30 * time.Second):
		fmt.Fprintf(c.Response(), "data: timeout\n\n")
	}
	flusher.Flush()

	return nil
}
