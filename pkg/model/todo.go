package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int64  `gorm:"primaryKey"`
	Task      string `gorm:"not null"`
	Completed bool   `gorm:"not null,default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Todo) TableName() string {
	return "todos"
}
