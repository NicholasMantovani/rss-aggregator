package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	businesslogic "github.com/NicholasMantovani/rssaggregator/internal/businessLogic"
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

	db := database.New(conn)

	apiConf := handlers.ApiConfig{DB: db}

	go businesslogic.StartScraping(db, 10, time.Minute)

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
	routerV1.Post("/feed_follows", apiConf.MiddlewareAuth(apiConf.HandleCreateFeedFollow))
	routerV1.Get("/feed_follows", apiConf.MiddlewareAuth(apiConf.HandleGetFeedsFollow))
	routerV1.Delete("/feed_follows/{id}", apiConf.MiddlewareAuth(apiConf.HandleDeleteFeedFollow))

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
