package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// GORM hooks
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	return nil
}

func (t *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	log.Println(t.UpdatedAt)
	return nil
}
