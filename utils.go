package main

import "github.com/gin-gonic/gin"

func sendError(context *gin.Context, input string, status int) {
	context.IndentedJSON(status, gin.H{"message": input})
}
