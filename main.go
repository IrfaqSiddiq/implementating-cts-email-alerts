package main

import (
	"log"
	"os"

	"cts-alerts/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file -> ", err)
	}
	r := routes.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
