package main

import (
	"my_website/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.Connect()
	database.DropTables(db)
}
