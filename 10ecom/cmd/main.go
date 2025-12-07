package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/whiteblueskyss/golang/internal/env"
)

func main() {

	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dns: env.GetString("GOOSE_DBSTRING", "host=localhost port=5434 user=postgres password=postgres dbname=ecomdb sslmode=disable"),
		},
	}

	// database setup
	slog.Info("Connecting to database...", "dns", cfg.db.dns)

	conn, err := pgx.Connect(ctx, cfg.db.dns)
	if err != nil {
		slog.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

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
