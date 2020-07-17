package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	r = mux.NewRouter()
)

func main() {
	r.Methods(http.MethodPost).Path("/login").HandlerFunc(Login)
	r.Methods(http.MethodPost).Path("/refresh").HandlerFunc(Refresh)

	r.Methods(http.MethodPost).Path("/logout").Handler(TokenAuthMiddleware(Logout()))

	log.Fatal(http.ListenAndServe(":8090", r))
}
