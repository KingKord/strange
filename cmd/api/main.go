package main

import (
	"database/sql"
	"fmt"
	_ "github.com/KingKord/strange/docs"
	"github.com/KingKord/strange/internal/helpers"
	"github.com/KingKord/strange/internal/repository/postgres"
	"github.com/KingKord/strange/internal/services"
	"os"
	"time"

	"github.com/KingKord/strange/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const port = "3333"

var counts int64

// @title strange API
// @version 1.0

// @host localhost:3333
// @BasePath /api/v1/

func main() {
	log.Println("Starting strange service")

	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to postgres")
	}

	helpers.MigrateUp()

	repo := postgres.NewPostgresRepo(conn)
	scheduleService := services.NewScheduleService(repo)
	newHandlers := handlers.NewHandlers(scheduleService)

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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}
		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing of for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
