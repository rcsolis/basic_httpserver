package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rcsolis/basic_httpserver/internal/model"
)

func CreateTask(ctx *gin.Context) {
	var task model.TaskRequest

	if err := ctx.ShouldBindJSON(&task); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()
	date := time.Now().UTC().Format(time.RFC3339)

	if _, ok := model.TaskDB[id]; ok {
		ctx.JSON(http.StatusConflict, gin.H{"error": "ID already exists"})
		return
	}

	model.TaskDB[id] = model.Task{
		ID:          id,
		Title:       task.Title,
		Description: task.Description,
		Date:        date,
		Done:        false,
	}

	ctx.JSON(http.StatusCreated, model.TaskDB[id])
}
