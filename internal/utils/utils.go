package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/godofprodev/todo-app-backend/internal/models"
)

func SendError(context *gin.Context, status int, input string) {
	context.IndentedJSON(status, gin.H{"message": input})
}

func Remove(slice []models.Task, s int) []models.Task {
	return append(slice[:s], slice[s+1:]...)
}
