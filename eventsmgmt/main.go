package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"booking-app/eventsmgmt/pkg/events"

	"github.com/gorilla/handlers"
)

func main() {

	// Wrap logs in specific format
	loggedServer := handlers.LoggingHandler(os.Stdout, events.Handlers())

	// Service port should be changed
	server := http.Server{
		Addr:         "127.0.0.1:9000",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      loggedServer,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
