package main

import "net/http"

func MiddlewareAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := IsTokenValid(r)
		if err != nil {
			ResponseError(w, ErrAuth)
			return
		}
		next.ServeHTTP(w, r)
	}
}
