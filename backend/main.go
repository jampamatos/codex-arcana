package main

import (
	"embed"
	"net/http"

	"github.com/rs/cors"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Initialize a simple HTTP servcer for /ping and /api/version in goroutine
	go func() {
		// Create new Mux router
		mux := http.NewServeMux()

		// Handle /ping endpoint, to return a simple "pong" response
		mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		// Handle /api/version endpoint, to return a simple JSON response
		mux.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"version": "0.1.0"}`))
		})

		// Campaign API routes
		mux.HandleFunc("/api/campaigns", campaignsHandler)
		mux.HandleFunc("/api/campaigns/", campaignHandler)

		// Apply CORS middleware to the mux server
		handler := cors.New(cors.Options{
			AllowedOrigins: []string{"*"}, // Allow all origins
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				http.MethodOptions,
			},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true, // Allow credentials
		}).Handler(mux)

		// Start the HTTP server on port 3000 with CORS enabled
		http.ListenAndServe(":3000", handler)
	}()

	// Initialize the database
	initDB()
	defer DB.Close() // Ensure the database connection is closed when the application exits

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "backend",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
