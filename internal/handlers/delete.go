package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rcsolis/basic_httpserver/internal/model"
)

func DeleteTask(ctx *gin.Context) {

	id := ctx.Param("id")
	if _, ok := model.TaskDB[id]; !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}

	delete(model.TaskDB, id)
}
