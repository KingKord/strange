package handlers

import (
	"fmt"
	"github.com/KingKord/strange/internal/repository/test_repo"
	"github.com/KingKord/strange/internal/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func getRoutes(handler Handlers) http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Route("/api/v1", func(mux chi.Router) {
		mux.Get("/", handler.Root)
		mux.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://localhost:%s/api/v1/swagger/doc.json", "3333")), //The url pointing to API definition
		))

		mux.Route("/schedule", func(router chi.Router) {
			router.Get("/day", handler.DaySchedule)
			router.Post("/reserve", handler.AssignMeet)
		})

	})
	return mux
}

func setupTest() http.Handler {
	repo := test_repo.NewTestRepo()
	serv := services.NewScheduleService(repo)
	handler := NewHandlers(serv)

	return getRoutes(handler)
}
