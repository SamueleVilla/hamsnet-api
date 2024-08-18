package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samuelevilla/hasnet-api/handlers"
)

type APIServer struct {
	http.Server
	logger *log.Logger
}

type APIServerParams struct {
	Host   string
	Port   string
	Logger *log.Logger
}

func NewAPIServer(params APIServerParams) *APIServer {
	router := chi.NewRouter()
	return &APIServer{
		Server: http.Server{
			Addr:    params.Host + ":" + params.Port,
			Handler: router,
		},
		logger: params.Logger,
	}
}
func (s *APIServer) registeregisterMiddlewares(router *chi.Mux) {
	router.Use(middleware.Logger)
}

func (s *APIServer) registerRoutes(router *chi.Mux) {
	handlers.NewPingHandler().RegisterRoutes(router)
	handlers.NewHamsterHandler().RegisterRoutes(router)
}

func (s *APIServer) Start() error {
	s.registeregisterMiddlewares(s.Handler.(*chi.Mux))
	s.registerRoutes(s.Handler.(*chi.Mux))

	s.logger.Printf("Starting server on %s", s.Addr)
	return s.ListenAndServe()
}
