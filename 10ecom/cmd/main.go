package main

import (
	"log/slog"
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

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)
	// mount and run in one line
	if err := api.run(api.mount()); err != nil {
		// log.Fatal(err)
		slog.Error("Error starting server", "error", err)

		os.Exit(1)
	}
}
