package actions

import (
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func FlashHandler(c buffalo.Context) error {
	taskID := c.Param("taskID")

	if result, exists := globalMessageStore.Get(taskID); exists {
		c.Flash().Add(result.Type, result.Content)
		globalMessageStore.Delete(taskID)
	}

	return c.Render(200, r.HTML("_flash.html"))
}
