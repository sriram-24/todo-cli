package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id         uuid.UUID
	Todo       string
	CreatedAt  string
	ModifiedAt string
}

func NewTodo(todo string) Todo {

	return Todo{
		Id:         uuid.New(),
		Todo:       todo,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		ModifiedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
