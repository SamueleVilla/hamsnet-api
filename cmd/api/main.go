package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samuelevilla/hasnet-api/internal/api"
	"github.com/samuelevilla/hasnet-api/internal/config"
	"github.com/samuelevilla/hasnet-api/internal/database"
	"github.com/samuelevilla/hasnet-api/internal/handlers"
	"github.com/samuelevilla/hasnet-api/internal/store"

	_ "github.com/lib/pq"
)

func main() {

	logger := log.New(log.Writer(), "[hamsnet-api] -- ", log.LstdFlags)
	addr := fmt.Sprintf("%s:%s", config.Env.SERVER_HOST, config.Env.SERVER_PORT)

	db, err := database.NewPsql(database.PsqlConfig{
		User:     config.Env.POSTRGRES_USER,
		Password: config.Env.POSTGRES_PASSWORD,
		Host:     config.Env.POSTGRES_HOST,
		Port:     config.Env.POSTGRES_PORT,
		DBName:   config.Env.POSTGRES_DBNAME,
	})
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	logger.Println("Connected to database")

	store := store.NewPsqlStore(db, logger)
	server := api.NewAPIServer(api.APIServerParams{
		Addr:   addr,
		Logger: logger,
		Handlers: []api.Handler{
			handlers.NewPingHandler(),
			handlers.NewHamsterHandler(store, config.Env.JWT_SECRET),
			handlers.NewAuthHandler(store),
		},
	})

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	/*
		We start the server in a goroutine so that we can listen for signals
	*/
	go func() {
		logger.Printf("Starting server at %s", addr)
		if err := server.ListenAndServe(); err != nil {
			logger.Fatalf("Error starting server: %v", err)
		}
	}()

	// Graceful shutdown release resources before shutting down
	<-signalCh

	logger.Println("Shutting down server...")
	server.Shutdown(context.Background())
}
