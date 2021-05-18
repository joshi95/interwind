package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joshi95/interwind/interwind/app/api/handlers"
)

func run() error {
	log := log.New(os.Stdout, "interwind", log.LUTC|log.Lshortfile|log.LstdFlags)

	app := handlers.API(log)

	api := http.Server{
		Addr:         ":8000",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("main : API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		log.Fatal(err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("main: server terminated with %v", err)
	}
}
