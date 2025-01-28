package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rcsolis/basic_httpserver/internal/handlers"
)

func UpdateRoutes(router *gin.Engine) {
	router.PUT("/task/:id", handlers.UpdateTask)
}
