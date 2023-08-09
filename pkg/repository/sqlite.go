package repository

import (
	"context"

	"github.com/ganeshagrawal55/sample-gin-gorm/pkg/model"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	db *gorm.DB
}

func NewSQLiteRepository(db *gorm.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (repo *SQLiteRepository) Migrate(ctx context.Context) error {
	models := []interface{}{
		&model.Todo{},
	}

	return repo.db.WithContext(ctx).AutoMigrate(models...)
}
