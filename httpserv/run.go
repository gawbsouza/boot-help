package httpserv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ShutdownTimeout = 5 * time.Second
)

func listenAndServe(server *http.Server) {
	log.Printf("Starting server at address %s\n", server.Addr)

	err := server.ListenAndServe()

	switch {
	case errors.Is(err, http.ErrServerClosed):
		log.Println("Server closed")
	default:
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func handleShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // Wait for some syscall errors

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server could not shutdown gracefully: %v\n", err)
	}

	log.Println("Server exiting")
}

// Listen and serve a http.Handler (for example, chi.NewRouter()) and handles server shutdown
func Run(port int, handler http.Handler) {
	addr := fmt.Sprintf(":%d", port)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	// Starts server in goroutine so that it won't block the shutdown handling below
	go listenAndServe(server)

	handleShutdown(server)
}
