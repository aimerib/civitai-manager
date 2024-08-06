package main

import (
	"civitai-manager/config"
	"civitai-manager/handlers"
	"civitai-manager/helpers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize handlers
	modelHandler := handlers.NewModelHandler(db)
	settingsHandler := handlers.NewSettingsHandler()
	utilHandler := handlers.NewUtilHandler()

	// Initialize template helpers
	templateHelpers := helpers.NewTemplateHelpers()

	r.SetFuncMap(templateHelpers.FuncMap())
	r.LoadHTMLGlob("templates/**/*")

	// Setup routes
	r.GET("/", modelHandler.ModelsIndex)
	r.GET("/models", modelHandler.ModelsIndex)
	r.GET("/models/:id", modelHandler.ModelsShow)
	r.GET("/settings", settingsHandler.Index)
	r.GET("/flash-partial/:taskID", utilHandler.FlashHandler)
	r.GET("/routes", utilHandler.RoutesHandler) // New handler for displaying routes

	// Example nested routes
	// r.GET("/models/:model_id/versions", modelHandler.VersionsIndex)
	// r.GET("/models/:model_id/versions/:id", modelHandler.VersionsShow)

	r.Static("/public", "./public")

	// Start the server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
