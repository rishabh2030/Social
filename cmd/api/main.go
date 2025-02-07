package main

import "log"

/*
	This is our entry point of our project
*/

func main() {

	cfg := config{
		addr: ":8080",
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))

}
