package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/samuelevilla/hasnet-api/internal/httputil"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) RegisterRoutes(router chi.Router) {
	router.Get("/ping", h.HandlePing) // GET http://localhost:8080/ping
}

// HandlePing godoc
// @Summary Responds with a pong message
// @Description Responds with a pong message
// @Tags ping
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func (h *PingHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "pong",
	}

	httputil.WriteJSON(w, http.StatusOK, data)
}
