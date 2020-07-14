package main

import (
	"github.com/gorilla/mux"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/go-openapi/runtime/middleware"

	"github.com/nicholasjackson/env"
	"github.com/machariamarigi/ju0920/product-api/handlers"
	"github.com/machariamarigi/ju0920/product-api/data"
)

var bindAddress = env.String("BindAddress", false, ":9090", "Bind Address To The Server")


func main()  {
	env.Parse()


	logger := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	validator := data.NewValidation()

	// create handlers
	productHandler := handlers.NewProducts(logger, validator)

	// create a new server multiplexer and register handlers with their routes
	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.ListSingle)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.Create)
	postRouter.Use(productHandler.MiddlewareProductValidation)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.Update)
	putRouter.Use(productHandler.MiddlewareProductValidation)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	docsHandler := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", docsHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// create a new server
	server := &http.Server {
		Addr:            ":9090",              // configure the bind address		
		Handler:         router,               // set the default handelr
		ErrorLog:        logger,               // set set the logger for the server
		IdleTimeout:     120 * time.Second,    // max time for connections using TCP Keep-Alive
		ReadTimeout:     1 * time.Second,      // max time to read request from client
		WriteTimeout:    1 * time.Second,      // max time to write request from client
	}


	// start the server
	go func() {
		logger.Println("Starting the server at port 9090")

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting serving: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt, gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <- sigChan
	logger.Println("Recieved terminate, graceful shutdown", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
