package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/samuelevilla/hasnet-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type APIServer struct {
	http.Server
}

type APIServerParams struct {
	Addr     string
	Logger   *log.Logger
	Handlers []Handler
}

// @title Hamsnet API
// @version 1.0
// @description This is a sample server Social Network server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host hamsnet.swagger.io
// @BasePath /api/v1
func NewAPIServer(params APIServerParams) *APIServer {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)

	// swagger endpoint
	swaggerEndpoint := fmt.Sprintf("http://%s/swagger/doc.json", params.Addr)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerEndpoint),
	))
	params.Logger.Printf("Swagger UI available at http://%s/swagger/index.html", params.Addr)

	// register routes
	for _, handler := range params.Handlers {
		handler.RegisterRoutes(router)
	}

	return &APIServer{
		Server: http.Server{
			Addr:    params.Addr,
			Handler: router,
		},
	}
}
