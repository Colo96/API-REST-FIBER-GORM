package models

import "gorm.io/gorm"

type Users struct {
	ID    uint    `gorm:"primary ; autoIncrement" json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
