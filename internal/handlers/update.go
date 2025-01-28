package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rcsolis/basic_httpserver/internal/model"
)

func UpdateTask(ctx *gin.Context) {

	var task model.TaskUpdateRequest
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, ok := model.TaskDB[id]; !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}
	model.TaskDB[id] = model.Task{
		ID:          id,
		Title:       task.Title,
		Description: task.Description,
		Date:        time.Now().UTC().Format(time.RFC3339),
		Done:        task.Done,
	}
	ctx.JSON(http.StatusOK, model.TaskDB[id])

}
