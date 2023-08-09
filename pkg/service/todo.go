package service

import (
	"context"

	"github.com/ganeshagrawal55/sample-gin-gorm/pkg/model"
	"github.com/ganeshagrawal55/sample-gin-gorm/pkg/repository"
)

type ITodoService interface {
	CreateNew(ctx context.Context, task string) (*model.Todo, error)
	GetByID(ctx context.Context, id int64) (*model.Todo, error)
	GetAll(ctx context.Context) ([]model.Todo, error)
	UpdateTask(ctx context.Context, id int64, task string) (*model.Todo, error)
	UpdateState(ctx context.Context, id int64, completed bool) (*model.Todo, error)
	DeleteById(ctx context.Context, id int64) error
}

type TodoService struct {
	repoTodo repository.ITodoRepository
}

func NewTodoService(repoTodo repository.ITodoRepository) ITodoService {
	return &TodoService{
		repoTodo: repoTodo,
	}
}

func (srv *TodoService) CreateNew(ctx context.Context, task string) (*model.Todo, error) {
	todo := model.Todo{
		Task:      task,
		Completed: false,
	}

	if err := srv.repoTodo.CreateTodo(ctx, &todo); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (srv *TodoService) GetByID(ctx context.Context, id int64) (*model.Todo, error) {
	todo, err := srv.repoTodo.GetTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (srv *TodoService) GetAll(ctx context.Context) ([]model.Todo, error) {
	todos, err := srv.repoTodo.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (srv *TodoService) UpdateTask(ctx context.Context, id int64, task string) (*model.Todo, error) {
	todo, err := srv.repoTodo.GetTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if todo == nil {
		return nil, nil
	}

	todo.Task = task
	if _, err := srv.repoTodo.UpdateTodo(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (srv *TodoService) UpdateState(ctx context.Context, id int64, completed bool) (*model.Todo, error) {
	todo, err := srv.repoTodo.GetTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if todo == nil {
		return nil, nil
	}

	todo.Completed = completed
	if _, err := srv.repoTodo.UpdateTodo(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (srv *TodoService) DeleteById(ctx context.Context, id int64) error {
	if _, err := srv.repoTodo.DeleteTodo(ctx, id); err != nil {
		return err
	}

	return nil
}
