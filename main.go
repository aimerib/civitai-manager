package main

import (
	"civitai-manager/config"
	"civitai-manager/handlers"
	"civitai-manager/middleware"
	"civitai-manager/views"
	"civitai-manager/workers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.DebugMode)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()

	// Use custom middleware
	r.Use(middleware.CacheMiddleware())

	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup session middleware
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	r.Use(sessions.Sessions("_civitai_manager_session", store))
	r.Use(middleware.SessionData())

	// Use DB Transaction middleware
	r.Use(middleware.DBTransactionMiddleware(db))

	// Initialize TemplateHelpers
	// templateHelpers := helpers.NewTemplateHelpers()

	// // Debug: Print out the keys in the FuncMap
	// funcMap := templateHelpers.FuncMap()

	// // Add the FuncMap to the default template functions
	// r.SetFuncMap(funcMap)

	// Initialize handlers
	modelHandler := handlers.NewModelHandler()
	// utilHandler := handlers.NewUtilHandler()
	websocketHandler := handlers.NewWebsocketHandler()

	// Setup routes
	r.GET("/", WithLayout("Models", modelHandler.ModelsIndex))
	r.GET("/models", modelHandler.ModelsIndex)
	r.GET("/models/:id", modelHandler.ModelsShow)
	// r.GET("/flash-partial/:taskID", utilHandler.FlashHandler)
	// r.GET("/routes", utilHandler.RoutesHandler)

	// Initialize in-memory worker
	worker := workers.NewWorker(10) // 10 concurrent workers
	worker.RegisterModelFetchWorker(db)
	worker.Start()
	defer worker.Stop()

	// Initialize handlers
	settingsHandler := handlers.NewSettingsHandler(worker)

	// Setup routes
	r.GET("/settings", settingsHandler.Index)
	r.POST("/settings/run-fetch-job", settingsHandler.StartBackgroundFetchJob)
	r.GET("/ws/:taskID", websocketHandler.WebSocketHandler)

	r.Static("/assets", "./public/assets")
	r.Static("/robots.txt", "./public/robots.txt")
	r.Static("/sw.js", "./public/assets/js/sw.js")

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("Starting server on :" + port)

	// Add the new CSRF middleware
	csrfMiddleware := csrf.Protect(
		[]byte("32-byte-long-auth-key"), // Replace with a secure key
		csrf.Secure(false),              // Set to true in production
		csrf.Path("/"),
	)

	// Wrap your router with the CSRF middleware
	http.ListenAndServe(":"+port, csrfMiddleware(r))
}

func RenderView(c *gin.Context, view templ.Component) {
	c.Set("content", view)
}

func WithLayout(t string, h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c)
		if !c.IsAborted() {
			content := c.MustGet("content").(templ.Component)
			csrfField := csrf.TemplateField(c.Request)
			templWrapper := fmt.Sprint("%v", csrfField)
			fmt.Println("Content:", content)
			fmt.Println("TemplWrapper:", templWrapper)
			views.Layout(t, content, templWrapper).Render(c, c.Writer)
		}
	}
}
