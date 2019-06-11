package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const version = "v1.0.0"

func main() {
	log.Println("Starting our helloworld app")

	// Health check
	http.HandleFunc("/health/check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	// Current version
	http.HandleFunc("/v", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, version)
	})

	// Application routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	port := os.Getenv("APP_PORT")
	s := http.Server{Addr: ":" + port}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	// https://gobyexample.com/signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	fmt.Println("Shutting down")
	s.Shutdown(context.Background())
}
