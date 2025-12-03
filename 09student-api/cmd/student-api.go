package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"student-api/internal/config"
	"syscall"
	"time"
)

func main() {
	// load configuration

	cfg := config.MustLoad() // calling MustLoad function from config package to load configuration

	//database setup
	// server routes

	router := http.NewServeMux() // creating a new HTTP request multiplexer (router)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Student API!"))
	})

	//HTTP server setup
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	// err := server.ListenAndServe()

	// if err != nil {
	// 	log.Fatalf("Server failed to start: %v", err)
	// }
	// //graceful shutdown
	// fmt.Println("Server is running on", cfg.HTTPServer.Addr)

	slog.Info("server started", slog.String("address", cfg.Addr)) // slog is a structured logger. Here we log an info message indicating that the server has started, along with the address it's running on.

	done := make(chan os.Signal, 1) // channel to receive OS signals for graceful shutdown. Here we create a channel named "done" that can receive OS signals. 1 is the buffer size of the channel.

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // signal.Notify registers the given channel to receive notifications of the specified signals. Here, we register the "done" channel to receive notifications for interrupt signals (like Ctrl+C) and termination signals. syscall.SIGINT and syscall.SIGTERM are constants representing these signals. Working process will send these signals to the channel when they occur.

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done // wait for a signal to be received

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // create a context with a timeout of 5 seconds for the shutdown process. context is used to manage deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes. Working process will use this context to ensure that the shutdown process does not exceed 5 seconds.

	defer cancel() // defer cancel() means that the cancel function will be called at the end of the surrounding function (main in this case). This ensures that any resources associated with the context are released when the main function exits.

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")

}
