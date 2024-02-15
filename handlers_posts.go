package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := cfg.DB.GetPosts(r.Context(), 15)
	if err != nil {
		log.Print(err)
		respondWithError(w, 500, "Something went wrong!")
		return
	}
	respondWithJSON(w, 200, posts)
}
