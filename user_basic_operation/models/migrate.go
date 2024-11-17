package models

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := MigrateBook(db); err != nil {
		return err
	}
	if err := MigrateUser(db); err != nil {
		return err
	}

	return nil
}
