package router

import (
	"github.com/gin-gonic/gin"
	"github.com/godofprodev/todo-app-backend/internal/models"
	"github.com/godofprodev/todo-app-backend/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

var tasks = []models.Task{
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

func getTasksHandler(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tasks)
}

func getTaskHandler(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		utils.SendError(context, http.StatusBadRequest, "Invalid Task id")
		return
	}

	for _, task := range tasks {
		if task.Id == id {
			context.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	utils.SendError(context, http.StatusNotFound, "Task not found")
}

func addTaskHandler(context *gin.Context) {
	var newTask models.Task

	err := context.BindJSON(&newTask)
	if err != nil {
		return
	}

	tasks = append(tasks, newTask)

	context.IndentedJSON(http.StatusCreated, tasks)
}

func updateTaskHandler(context *gin.Context) {
	var rTask models.Task
	idStr := context.Param("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		utils.SendError(context, http.StatusBadRequest, "Invalid Task id")
		return
	}

	var newTask models.Task

	err = context.BindJSON(&newTask)
	if err != nil {
		return
	}

	if !newTask.Validate() {
		utils.SendError(context, http.StatusBadRequest, "Invalid Task data")
		return
	}

	for i := 0; i < len(tasks); i++ {
		task := &tasks[i]

		if task.Id == id {
			task.Title = newTask.Title
			task.Description = newTask.Description
			task.DueDate = newTask.DueDate
			task.Status = newTask.Status
			rTask = *task

			return
		}
	}

	if !rTask.Validate() {
		utils.SendError(context, http.StatusNotFound, "Task not found")
		return
	}

	context.IndentedJSON(http.StatusOK, rTask)
}

func patchTaskHandler(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		utils.SendError(context, http.StatusBadRequest, "Invalid Task id")
		return
	}

	var input models.Task

	err = context.BindJSON(&input)
	if err != nil {
		utils.SendError(context, http.StatusBadRequest, "Invalid Task data")
		return
	}

	for i := 0; i < len(tasks); i++ {
		task := &tasks[i]

		if task.Id == id {
			task.UpdateTask(input)
			context.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	utils.SendError(context, http.StatusNotFound, "Task not found")
}

func deleteTaskHandler(context *gin.Context) {
	var rTask models.Task
	idStr := context.Param("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		utils.SendError(context, http.StatusBadRequest, "Invalid Task id")
		return
	}

	indexToRemove := -1

	for i, task := range tasks {
		if task.Id == id {
			rTask = task
			indexToRemove = i
			return
		}
	}

	if !rTask.Validate() {
		utils.SendError(context, http.StatusNotFound, "Task not found")
		return
	}

	tasks = utils.Remove(tasks, indexToRemove)

	context.IndentedJSON(http.StatusOK, tasks)
}
