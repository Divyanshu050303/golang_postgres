package models

import "gorm.io/gorm"

type UserModels struct {
	ID           string `gorm:"type uuid; primary key" json:"id"`
	UserName     string `json:"userName"`
	UserEmail    string `json:"userEmail"`
	UserPassword string `json:"userPassword" validate:"min=8"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&UserModels{})
}
