package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/NicholasMantovani/rssaggregator/internal/database"
	"github.com/NicholasMantovani/rssaggregator/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No PORT specified")
	}

	dbUrl := os.Getenv("DB_URL")
	if port == "" {
		log.Fatal("No DB_URL specified")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	apiConf := handlers.ApiConfig{DB: database.New(conn)}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routerV1 := chi.NewRouter()
	routerV1.Get("/healthz", handlers.Readiness)
	routerV1.Get("/err", handlers.Error)
	routerV1.Post("/users", apiConf.HandleCreateUser)
	routerV1.Get("/users", apiConf.MiddlewareAuth(apiConf.HandleGetUser))
	routerV1.Post("/feeds", apiConf.MiddlewareAuth(apiConf.HandleCreateFeed))
	routerV1.Get("/feeds", apiConf.HandleGetFeeds)

	router.Mount("/v1", routerV1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Print("Application is availabe on port: ", port)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
