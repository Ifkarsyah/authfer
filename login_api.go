package main

import (
	"encoding/json"
	"net/http"
)

type loginHandlerFunc func(u *User) (*TokenDetails, error)

func LoginAPI(loginHandler loginHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := new(User)

		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ts, err := loginHandler(u)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ResponseOK(w, ts)
	})
}
