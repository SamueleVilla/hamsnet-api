package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/samuelevilla/hasnet-api/internal/store"
)

type AuthHandler struct {
	store *store.Store
}

func NewAuthHandler(store *store.Store) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (h *AuthHandler) RegisterRoutes(router chi.Router) {
	router.Post("/login", h.Login)
	router.Post("/register", h.Register)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
}
