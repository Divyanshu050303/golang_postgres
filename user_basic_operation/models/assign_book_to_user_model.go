package models

import (
	"time"

	"gorm.io/gorm"
)

type AssignBookToUserModel struct {
	UserID     string    `gorm:"type:uuid not null" json:"userId"`
	BookID     string    `gorm:"type:uuid not null" json:"bookId"`
	AssignedAt time.Time `json:"assignedAt"`
	ReturnAt   time.Time `json:"returnAt"`
}

func AssignBookToUser(db *gorm.DB) error {
	err := db.AutoMigrate(&AssignBookToUserModel{})
	return err
}
