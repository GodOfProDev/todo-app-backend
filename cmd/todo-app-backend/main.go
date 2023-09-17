package main

import (
	"github.com/godofprodev/todo-app-backend/internal/router"
)

func main() {
	router := router.NewRouter()

	err := router.Start()

	if err != nil {
		panic(err)
	}
}
