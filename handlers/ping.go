package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) RegisterRoutes(router *chi.Mux) {
	router.Get("/ping", h.HandlePing) // GET http://localhost:8080/ping
}

func (h *PingHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := map[string]string{
		"message": "pong",
		"status":  "ok",
	}

	json.NewEncoder(w).Encode(data)
}
