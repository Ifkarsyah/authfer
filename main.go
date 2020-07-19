package main

import (
	"fmt"
	"github.com/Ifkarsyah/authfer/api"
	"github.com/Ifkarsyah/authfer/pkg/config"
	"github.com/Ifkarsyah/authfer/repo"
	"github.com/Ifkarsyah/authfer/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func NewRouter(svc service.Service) *mux.Router {
	root := mux.NewRouter()

	root.Methods(http.MethodPost).Path("/login").Handler(api.LoginAPI(svc.Login))
	root.Methods(http.MethodPost).Path("/refresh").Handler(api.MiddlewareAuth(api.RefreshAPI(svc.RefreshToken)))
	root.Methods(http.MethodPost).Path("/logout").Handler(api.MiddlewareAuth(api.Logout(svc.Logout)))

	return root
}

func NewService(dep *Dependency) service.Service {
	return service.Service{
		Redis: dep.cache,
		DB:    dep.db,
	}
}

type Dependency struct {
	cache *repo.RedisRepo
	db    *repo.DBRepo
}

func main() {
	svc := NewService(&Dependency{
		cache: repo.NewRedisConnection(config.AppConfig.RedisHost, config.AppConfig.RedisPort),
		db:    repo.NewDBConnection(),
	})

	server := &http.Server{
		Handler:      NewRouter(svc),
		Addr:         fmt.Sprintf("%s:%s", config.AppConfig.Host, config.AppConfig.Port),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
