package main

import (
	"log"

	"github.com/damuel90/twitgo/db"
	"github.com/damuel90/twitgo/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if !db.CheckConnection() {
		log.Fatal("Database connection failed")
		return
	}
	handlers.RunHandlers()
}
