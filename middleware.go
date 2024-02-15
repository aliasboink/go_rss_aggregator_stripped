package main

import (
	"log"
	"net/http"
	"strings"
)

type authedHandler func(http.ResponseWriter, *http.Request)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apikeyString := strings.TrimPrefix(r.Header.Get("Authorization"), "ApiKey ")
		log.Println("Haha", apikeyString)
		log.Println("Haha2", cfg.ApiKey)
		if apikeyString != cfg.ApiKey {
			respondWithError(w, 401, "Unauthorized!")
			return
		}
		handler(w, r)
	})
}
