package main

import (
	"github.com/joho/godotenv"
	"log"
)

/*
	This is our entry point of our project
*/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	cfg := config{
		addr: getString("DEV_PORT", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))

}
