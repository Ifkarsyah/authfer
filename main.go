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
	r.Methods(http.MethodPost).Path("/login").Handler(LoginAPI(LoginHandler))
	r.Methods(http.MethodPost).Path("/refresh").Handler(RefreshAPI())
	r.Methods(http.MethodPost).Path("/logout").Handler(MiddlewareAuth(Logout()))

	log.Fatal(http.ListenAndServe(":8090", r))
}
