package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rcsolis/basic_httpserver/internal/routes"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()

	router := gin.Default()
	routes.CreateRoutes(router)
	routes.UpdateRoutes(router)
	routes.DeleteRoutes(router)
	routes.ListRoutes(router)

	s := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
