package main

import (
	"log"
	"net/http"
	"rss/internal/database"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request) {
	offset := chi.URLParam(r, "offset")
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		respondWithError(w, 500, "Something went wrong!")
		return
	}
	posts, err := cfg.DB.GetPosts(r.Context(), database.GetPostsParams{
		Limit:  15,
		Offset: int32(offsetInt),
	})
	if err != nil {
		log.Print(err)
		respondWithError(w, 500, "Something went wrong!")
		return
	}
	respondWithJSON(w, 200, posts)
}
