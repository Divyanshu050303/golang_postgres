package repository

import "gorm.io/gorm"

type BookRepository struct {
	DB *gorm.DB
}
