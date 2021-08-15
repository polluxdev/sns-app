package main

import (
	"log"
	"sns-app/config"
	"sns-app/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	config.DBInit()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
