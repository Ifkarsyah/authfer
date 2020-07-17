package main

import (
	"net/http"
)

func Logout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		au, err := ExtractTokenMetadata(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		deleted, delErr := RedisDeleteAuth(au.AccessUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
