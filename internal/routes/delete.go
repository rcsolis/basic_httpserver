package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rcsolis/basic_httpserver/internal/handlers"
)

func DeleteRoutes(router *gin.Engine) {
	router.DELETE("/task/:id", handlers.DeleteTask)
}
