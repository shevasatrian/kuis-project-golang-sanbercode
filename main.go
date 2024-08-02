package main

import (
	"book-category-api/cmd/api"
	"book-category-api/internal/models"
	"log"
	"os"
)

func main() {
	dataSourceName := os.Getenv("DATABASE_URL")
	if dataSourceName == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	if err := models.InitDB(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	router := api.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.")
	}
}
