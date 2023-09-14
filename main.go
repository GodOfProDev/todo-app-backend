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

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTask)
	router.POST("/tasks", addTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run("localhost:8080")
}

func getTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tasks)
}

func getTask(context *gin.Context) {
	var rTask task
	idStr := context.Param("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, "Invalid task id")
		return
	}

	for _, task := range tasks {
		if task.Id != id {
			continue
		}

		rTask = task
	}

	if rTask.Status == "" {
		context.IndentedJSON(http.StatusNotFound, "Task not found")
		return
	}

	context.IndentedJSON(http.StatusOK, rTask)
}

func addTask(context *gin.Context) {
	var newTask task

	err := context.BindJSON(&newTask)
	if err != nil {
		return
	}

	tasks = append(tasks, newTask)

	context.IndentedJSON(http.StatusCreated, tasks)
}

func updateTask(context *gin.Context) {
	var rTask task
	idStr := context.Param("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, "Invalid task id")
		return
	}

	var newTask task

	err = context.BindJSON(&newTask)
	if err != nil {
		return
	}

	if newTask.Status == "" {
		context.IndentedJSON(http.StatusBadRequest, "Invalid task data")
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

		println(task.Title)
	}

	if rTask.Status == "" {
		context.IndentedJSON(http.StatusNotFound, "Task not found")
		return
	}

	context.IndentedJSON(http.StatusOK, rTask)
}

func deleteTask(context *gin.Context) {
	var rTask task
	idStr := context.Param("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, "Invalid task id")
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

	if rTask.Status == "" {
		context.IndentedJSON(http.StatusNotFound, "Task not found")
		return
	}

	tasks = remove(tasks, indexToRemove)

	context.IndentedJSON(http.StatusOK, tasks)
}

func remove(slice []task, s int) []task {
	return append(slice[:s], slice[s+1:]...)
}
