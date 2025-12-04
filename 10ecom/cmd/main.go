package main

import (
	"log"
	"os"
)

func main() {

	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := &application{
		config: cfg,
	}

	// h := api.mount()

	// err := api.run(h)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mount and run in one line
	if err := api.run(api.mount()); err != nil {
		// log.Fatal(err)
		log.Printf("Error starting server: %v", err)
		os.Exit(1)
	}
}
