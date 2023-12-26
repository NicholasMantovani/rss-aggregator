package main

import (
	"log"
	"net/http"
	"os"

	"github.com/NicholasMantovani/rssaggregator/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No PORT specified")
	}

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

	router.Mount("/v1", routerV1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Print("Application is availabe on port: ", port)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
