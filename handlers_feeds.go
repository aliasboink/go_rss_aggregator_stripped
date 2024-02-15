package main

import (
	"encoding/json"
	"log"
	"net/http"
	"rss/internal/database"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerPostFeed(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Print(err)
		respondWithError(w, 500, "Something went wrong!")
		return
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
	})
	if err != nil {
		log.Print(err)
		respondWithError(w, 500, "Something went wrong!")
		return
	}
	respondWithJSON(w, 200, feed)
}

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		log.Print(err)
		respondWithError(w, 500, "Something went wrong!")
		return
	}
	respondWithJSON(w, 200, feeds)

}
