package main

import (
	"github.com/godofprodev/todo-app-backend/internal/router"
)

func main() {
	apiRouter := router.NewRouter()

	err := apiRouter.Start()

	if err != nil {
		panic(err)
	}
}
