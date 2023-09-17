package router

import "github.com/gin-gonic/gin"

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Start() error {
	router := gin.Default()

	router.GET("/tasks", getTasksHandler)
	router.GET("/tasks/:id", getTaskHandler)
	router.POST("/tasks", addTaskHandler)
	router.PUT("/tasks/:id", updateTaskHandler)
	router.PATCH("/tasks/:id", patchTaskHandler)
	router.DELETE("/tasks/:id", deleteTaskHandler)

	return router.Run(":8080")
}
