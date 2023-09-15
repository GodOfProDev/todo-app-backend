package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     int64  `json:"due_date"`
	Status      string `json:"status"`
}

func (t task) validate() bool {
	return t.Title != "" && t.Description != "" && t.Status != ""
}

func (t task) validateTitle() bool {
	return t.Title != ""
}

func (t task) validateDescription() bool {
	return t.Description != ""
}

func (t task) validateStatus() bool {
	return t.Status != ""
}

func (t task) validateDueDate() bool {
	return t.DueDate != 0
}

var tasks = []task{
	{
		Id:          0,
		Title:       "Dishes",
		Description: "Clean the dishes",
		DueDate:     1694715520395,
		Status:      "Incomplete",
	},
	{
		Id:          1,
		Title:       "Dinner",
		Description: "Eat dinner",
		DueDate:     1694715620395,
		Status:      "Complete",
	},
	{
		Id:          3,
		Title:       "Gym",
		Description: "Go to the gym",
		DueDate:     1694715521395,
		Status:      "Incomplete",
	},
}

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasksRoute)
	router.GET("/tasks/:id", getTaskRoute)
	router.POST("/tasks", addTaskRoute)
	router.PUT("/tasks/:id", updateTaskRoute)
	router.PATCH("/tasks/:id", patchTaskRoute)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run("localhost:8080")
}

func getTasksRoute(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tasks)
}

func getTaskRoute(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task id"})
		return
	}

	for _, task := range tasks {
		if task.Id == id {
			context.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func addTaskRoute(context *gin.Context) {
	var newTask task

	err := context.BindJSON(&newTask)
	if err != nil {
		return
	}

	tasks = append(tasks, newTask)

	context.IndentedJSON(http.StatusCreated, tasks)
}

func updateTaskRoute(context *gin.Context) {
	var rTask task
	idStr := context.Param("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task id"})
		return
	}

	var newTask task

	err = context.BindJSON(&newTask)
	if err != nil {
		return
	}

	if !newTask.validate() {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task data"})
		return
	}

	for i := 0; i < len(tasks); i++ {
		task := &tasks[i]

		if task.Id != id {
			continue
		}

		task.Title = newTask.Title
		task.Description = newTask.Description
		task.DueDate = newTask.DueDate
		task.Status = newTask.Status
		rTask = *task
	}

	if !rTask.validate() {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, rTask)
}

func patchTaskRoute(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task id"})
		return
	}

	var input task

	err = context.BindJSON(&input)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task data"})
		return
	}

	for i := 0; i < len(tasks); i++ {
		task := &tasks[i]

		if task.Id == id {
			updateTask(input, task)
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func updateTask(input task, task *task) {
	if input.validateTitle() {
		task.Title = input.Title
	}
	if input.validateDescription() {
		task.Description = input.Description
	}
	if input.validateStatus() {
		task.Status = input.Status
	}
	if input.validateDueDate() {
		task.DueDate = input.DueDate
	}
}

func deleteTask(context *gin.Context) {
	var rTask task
	idStr := context.Param("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task id"})
		return
	}

	indexToRemove := -1

	for i, task := range tasks {
		if task.Id != id {
			continue
		}

		rTask = task
		indexToRemove = i
	}

	if !rTask.validate() {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	tasks = remove(tasks, indexToRemove)

	context.IndentedJSON(http.StatusOK, tasks)
}

func remove(slice []task, s int) []task {
	return append(slice[:s], slice[s+1:]...)
}
