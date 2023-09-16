package main

import "github.com/gin-gonic/gin"

func sendError(context *gin.Context, status int, input string) {
	context.IndentedJSON(status, gin.H{"message": input})
}
