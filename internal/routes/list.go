package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rcsolis/basic_httpserver/internal/handlers"
)

func ListRoutes(router *gin.Engine) {
	group := router.Group("/tasks")
	group.GET("/:id", handlers.GetTask)
	group.GET("/", handlers.ListTasks)
}
