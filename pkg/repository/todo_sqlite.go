package repository

import (
	"context"
	"errors"

	"github.com/ganeshagrawal55/sample-gin-gorm/pkg/model"
	"gorm.io/gorm"
)

func (repo *SQLiteRepository) CreateTodo(ctx context.Context, todo *model.Todo) error {
	if err := repo.db.WithContext(ctx).Create(todo).Error; err != nil {
		return err
	}

	return nil
}

func (repo *SQLiteRepository) GetAllTodos(ctx context.Context) ([]model.Todo, error) {
	var todo []model.Todo
	if err := repo.db.WithContext(ctx).Find(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return todo, nil
}

func (repo *SQLiteRepository) GetTodoByID(ctx context.Context, id int64) (*model.Todo, error) {
	var todo model.Todo
	if err := repo.db.WithContext(ctx).First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &todo, nil
}

func (repo *SQLiteRepository) UpdateTodo(ctx context.Context, todo *model.Todo) (bool, error) {
	result := repo.db.WithContext(ctx).Save(todo)
	if err := result.Error; err != nil {
		return false, err
	}

	if result.RowsAffected != 1 {
		return false, nil
	}

	return true, nil
}

func (repo *SQLiteRepository) DeleteTodo(ctx context.Context, id int64) (bool, error) {
	result := repo.db.WithContext(ctx).Delete(&model.Todo{}, id)
	if err := result.Error; err != nil {
		return false, err
	}

	if result.RowsAffected != 1 {
		return false, nil
	}

	return true, nil
}
