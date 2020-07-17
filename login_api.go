package main

import (
	"encoding/json"
	"net/http"
)

type loginHandlerFunc func(*LoginParams) (*LoginOutput, error)

func LoginAPI(loginHandler loginHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := new(LoginParams)
		err := json.NewDecoder(r.Body).Decode(params)
		if err != nil {
			ResponseError(w, ErrBadRequest)
			return
		}

		ts, err := loginHandler(params)
		if err != nil {
			ResponseError(w, err)
			return
		}

		ResponseOK(w, ts)
	})
}
