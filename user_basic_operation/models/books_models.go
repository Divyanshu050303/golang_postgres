package models

import "gorm.io/gorm"

type BookModel struct {
	ID          string  `gorm:"type:uuid; primary key" json:"id"`
	Author      *string `json:"author"`
	Publication *string `json:"publication"`
	Title       *string `json:"title"`
}

func MigrateBook(db *gorm.DB) error {
	err := db.AutoMigrate(&BookModel{})
	return err
}
