// server/main.go
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/martbul/locNot/api"
)

func main() {
	logger := hclog.Default()

	// Initialize the server with the configured router and logger
	server := api.NewServer(logger)

	// Set up the HTTP server
	srv := http.Server{
		Addr:         ":9000",
		Handler:      server.Router,
		ErrorLog:     logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Starting API server on port 9000")
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown handling
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, os.Kill)
	sig := <-signalChannel
	logger.Info("Received signal, shutting down", "signal", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(timeoutContext)
}
