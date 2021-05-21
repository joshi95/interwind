package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joshi95/interwind/interwind/app/interwind-api/handlers"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("main: server terminated with %v", err)
	}
}

func run() error {

	log := log.New(os.Stdout, "interwind", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	app := handlers.API(log)

	api := http.Server{
		Addr:         ":8000",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("main : api listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatal(err)
	case <-shutdown:
		log.Println("main: starting load shed")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := api.Shutdown(ctx)
		if err != nil {
			log.Printf("main : graceful shutdown did not complete")
			err = api.Close()
		}

		if err != nil {
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
