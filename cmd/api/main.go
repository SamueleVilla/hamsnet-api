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
	env := config.InitEnviroment()

	logger := log.New(log.Writer(), "[hamsnet-api] -- ", log.LstdFlags)
	addr := fmt.Sprintf("%s:%s", env.SERVER_HOST, env.SERVER_PORT)

	db, err := database.NewPsql(database.PsqlConfig{
		User:     env.POSTRGRES_USER,
		Password: env.POSTGRES_PASSWORD,
		Host:     env.POSTGRES_HOST,
		Port:     env.POSTGRES_PORT,
		DBName:   env.POSTGRES_DBNAME,
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
			handlers.NewHamsterHandler(store, env.JWT_SECRET),
			handlers.NewAuthHandler(store, env.JWT_SECRET),
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
