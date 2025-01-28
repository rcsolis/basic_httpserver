package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rcsolis/basic_httpserver/internal/model"
)

func ListTasks(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, model.TaskDB)
}

func GetTask(ctx *gin.Context) {

	id := ctx.Param("id")
	if task, ok := model.TaskDB[id]; ok {
		ctx.JSON(http.StatusOK, task)
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
}
