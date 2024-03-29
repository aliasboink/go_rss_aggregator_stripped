package main

import (
	"net/http"
	"strings"
)

type authedHandler func(http.ResponseWriter, *http.Request)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apikeyString := strings.TrimPrefix(r.Header.Get("Authorization"), "ApiKey ")
		if apikeyString != cfg.ApiKey {
			respondWithError(w, 401, "Unauthorized!")
			return
		}
		handler(w, r)
	})
}
