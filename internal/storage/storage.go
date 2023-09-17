package storage

import "github.com/godofprodev/todo-app-backend/internal/models"

type Storage interface {
	GetUserById(id int) *models.User
}
