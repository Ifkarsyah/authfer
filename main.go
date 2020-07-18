package main

import (
	"github.com/Ifkarsyah/authfer/api"
	"github.com/Ifkarsyah/authfer/handler"
	"github.com/Ifkarsyah/authfer/pkg/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	r = mux.NewRouter()
)

func main() {
	config.InitAppConfig()

	r.Methods(http.MethodPost).Path("/login").Handler(api.LoginAPI(handler.LoginHandler))
	r.Methods(http.MethodPost).Path("/refresh").Handler(api.RefreshAPI())
	r.Methods(http.MethodPost).Path("/logout").Handler(api.MiddlewareAuth(api.Logout()))

	log.Fatal(http.ListenAndServe(":8090", r))
}
