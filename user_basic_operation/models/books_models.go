package models

import "gorm.io/gorm"

type BookModels struct {
	ID          string  `gorm:"type:uuid;primaryKey" json:"id"`
	Author      *string `json:"author"`
	Publication string  `json:"publication"`
	Title       *string `json:"title"`
}

func MigrateBook(db *gorm.DB) error {
	err := db.AutoMigrate(&BookModels{})
	return err
}
func (BookModels) TableName() string {
	return "book_models"
}
