package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log"
)

func main() {
	log.Println("Starting server...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request:", r.Method, r.URL)
		fmt.Fprintln(w, "Hello Container!!!")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Could not start server: %v\n", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	fmt.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Could not gracefully shutdown the server: %v\n", err)
	}

	fmt.Println("Server stopped")
}