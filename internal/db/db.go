package db

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

func New(db gorm.Dialector) (*gorm.DB, error) {
	return gorm.Open(db, &gorm.Config{})
}

func PostgresDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_PORT"))
}
