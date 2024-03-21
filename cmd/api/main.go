package main

import (
	"github.com/KingKord/strange/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func main() {
	newHandlers := handlers.NewHandlers()

	log.Println("Starting service on port 3333")
	err := http.ListenAndServe(":3333", getRoutes(newHandlers))
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

	mux.Get("/", handler.Root)

	mux.Route("/schedule", func(router chi.Router) {
		router.Get("/day", handler.DaySchedule)
		router.Post("/reserve", handler.AssignMeet)
	})

	return mux
}
