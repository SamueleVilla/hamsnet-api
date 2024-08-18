package main

import (
	"log"

	"github.com/samuelevilla/hasnet-api/internal/api"
	"github.com/samuelevilla/hasnet-api/internal/config"
)

func main() {

	logger := log.New(log.Writer(), "hamsnet-api", log.LstdFlags)

	server := api.NewAPIServer(api.APIServerParams{
		Host:   config.Env.SERVER_HOST,
		Port:   config.Env.SERVER_PORT,
		Logger: logger,
	})

	server.Start()

	logger := log.New(log.Writer(), "hamsnet-api", log.LstdFlags)

	server := api.NewAPIServer(api.APIServerParams{
		Host:   config.Env.SERVER_HOST,
		Port:   config.Env.SERVER_PORT,
		Logger: logger,
	})

	server.Start()
}
