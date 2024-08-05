package actions

import (
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/worker"
	"github.com/google/uuid"
)

func generateUniqueID() string {
	return uuid.New().String()
}

// StartBackgroundTask initiates the long-running task
func StartBackgroundFetchJob(c buffalo.Context) error {
	taskID := generateUniqueID()
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		limit = 0 // Default value if not provided or invalid
	}

	err = w.Perform(worker.Job{
		Queue:   "default",
		Handler: "fetch_models",
		Args: worker.Args{
			"taskID": taskID,
			"limit":  limit,
		},
	})

	if err != nil {
		return c.Error(500, err)
	}

	return c.Render(200, r.JSON(map[string]string{"taskID": taskID}))
}
