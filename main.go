package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/nilerajput91/userloginapp/controller"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := controller.App{}
	app.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	app.RunServer()
}
