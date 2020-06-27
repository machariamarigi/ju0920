package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/machariamarigi/ju0920/product-api/handlers"
)

func main()  {
	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)

	// create handlers
	ph := handlers.NewProducts(l)

	// create a new server multiplexer and register handlers with their routes
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new server
	s := &http.Server {
		Addr:            ":9090",              // configure the bind address		
		Handler:         sm,                   // set the default handelr
		ErrorLog:        l,                    // set set the logger for the server
		IdleTimeout:     120 * time.Second,    // max time for connections using TCP Keep-Alive
		ReadTimeout:     1 * time.Second,      // max time to read request from client
		WriteTimeout:    1 * time.Second,      // max time to write request from client
	}


	// start the server
	go func() {
		l.Println("Starting the server at port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting serving: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt, gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
