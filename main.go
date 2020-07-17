package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", Login)
	router.POST("/token/refresh", Refresh)

	authorized := router.Group("/", TokenAuthMiddleware())
	{
		authorized.POST("/logout", Logout)
	}

	log.Fatal(router.Run(":8090"))
}
