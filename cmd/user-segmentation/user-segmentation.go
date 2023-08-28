package main

import (
	"log"

	apihandler "github.com/XineAurora/user-segmentation/internal/api-handler"
	"github.com/XineAurora/user-segmentation/internal/db"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
)

func main() {
	router := gin.Default()
	db, err := db.New(postgres.Open(db.PostgresDsn()))
	if err != nil {
		//TODO: add description to error
		log.Fatal(err)
	}
	handler, err := apihandler.New(router, db)

	handler.Start()
}
