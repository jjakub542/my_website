package main

import (
	"log"
	"my_website/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.Connect()
	err := database.InsertExampleArticles(db)
	if err != nil {
		log.Fatal(err)
	}
}
