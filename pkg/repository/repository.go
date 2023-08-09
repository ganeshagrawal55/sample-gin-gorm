package repository

import (
	"context"

	"github.com/ganeshagrawal55/sample-gin-gorm/pkg/model"
)

type ITodoRepository interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
	GetAllTodos(ctx context.Context) ([]model.Todo, error)
	GetTodoByID(ctx context.Context, id int64) (*model.Todo, error)
	UpdateTodo(ctx context.Context, todo *model.Todo) (bool, error)
	DeleteTodo(ctx context.Context, id int64) (bool, error)
}
