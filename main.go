package main

import (
	"civitai-manager/config"
	"civitai-manager/handlers"
	"civitai-manager/helpers"
	"civitai-manager/middleware"
	"html/template"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := gin.Default()

	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup session middleware
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	r.Use(sessions.Sessions("_civitai_manager_session", store))
	r.Use(middleware.SessionData())

	// Use the parameter logger middleware
	r.Use(middleware.ParamLogger(logger))

	// Use CSRF middleware
	r.Use(middleware.CSRF())

	// Use DB Transaction middleware
	r.Use(middleware.DBTransactionMiddleware(db))

	// Initialize template helpers
	templateHelpers := helpers.NewTemplateHelpers()

	r.SetFuncMap(templateHelpers.FuncMap())
	r.SetFuncMap(template.FuncMap{
		"csrfField": func(c *gin.Context) template.HTML {
			token := middleware.GetCSRFToken(c)
			return template.HTML(`<input type="hidden" name="csrf_token" value="` + token + `">`)
		},
	})

	r.LoadHTMLGlob("templates/**/*")

	// Initialize handlers
	modelHandler := handlers.NewModelHandler()
	settingsHandler := handlers.NewSettingsHandler()
	utilHandler := handlers.NewUtilHandler()

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
