package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"rss/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB     *database.Queries
	ApiKey string
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment!")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("dbURL is not found in the environment!")
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("dbURL is not found in the environment!")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB:     dbQueries,
		ApiKey: apiKey,
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerPostFeed))
	// v1router.Delete("/feeds", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeed))
	v1router.Get("/readiness", handlerReadiness)
	v1router.Get("/err", handlerError)
	v1router.Get("/posts", apiCfg.handlerGetPosts)

	r.Mount("/v1", v1router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go rssThiefWorker(dbQueries, 10, 10)

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
