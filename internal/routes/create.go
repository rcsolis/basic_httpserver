package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rcsolis/basic_httpserver/internal/handlers"
)

func CreateRoutes(router *gin.Engine) {
	router.POST("/task", handlers.CreateTask)
}
