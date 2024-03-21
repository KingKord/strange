package main

import (
	"fmt"
	_ "github.com/KingKord/strange/docs"

	"github.com/KingKord/strange/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

const port = "3333"

// @title strange API
// @version 1.0

// @host localhost:3333
// @BasePath /api/v1/

func main() {
	newHandlers := handlers.NewHandlers()

	log.Println(fmt.Sprintf("Starting service on port %s", port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), getRoutes(newHandlers))
	if err != nil {
		log.Fatalf("cannot start the application due: %s", err)
	}
}

func getRoutes(handler handlers.Handlers) http.Handler {
	mux := chi.NewMux()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)), //The url pointing to API definition
	))

	mux.Route("/api/v1", func(mux chi.Router) {
		mux.Get("/", handler.Root)

		mux.Route("/schedule", func(router chi.Router) {
			router.Get("/day", handler.DaySchedule)
			router.Post("/reserve", handler.AssignMeet)
		})

	})

	return mux
}
