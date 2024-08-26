package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	router.Use(middleware.ContentCharset("utf-8"))

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
