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

func NewRouter(handler service.Service) *mux.Router {
	root := mux.NewRouter()

	root.Methods(http.MethodPost).Path("/login").Handler(api.LoginAPI(handler.Login))
	//root.Methods(http.MethodPost).Path("/refresh").Service(api.RefreshAPI())
	root.Methods(http.MethodPost).Path("/logout").Handler(api.MiddlewareAuth(api.Logout(handler.Logout)))

	return root
}

func NewHandler(dep *Dependency) service.Service {
	return service.Service{
		Cacher: dep.cache,
	}
}

type Dependency struct {
	cache *repo.RedisRepo
}

func main() {
	c := config.InitAppConfig()

	h := NewHandler(&Dependency{
		cache: repo.NewRedisConnection(c.RedisHost, c.RedisPort),
	})

	srv := &http.Server{
		Handler:      NewRouter(h),
		Addr:         fmt.Sprintf("%s:%s", c.Host, c.Port),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
