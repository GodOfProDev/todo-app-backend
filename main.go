package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasks)

	router.Run("localhost:8080")
}

func getTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "hello")
}
