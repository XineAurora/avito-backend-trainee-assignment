package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(""), &gorm.Config{})
}
