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
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
	})
}

// HandleLogin godoc
// @Summary Responds with user id
// @Description Responds with user id
// @Tags auth
// @Produce json
// @FormParam username string true "Username"
// @FormParam password string true "Password"
// @Success 200 {object} types.AuthUserResponse
// @Failure 400 {object} httputil.HttpError
// @Failure 500 {object} httputil.HttpError
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
}

// HandleLogin godoc
// @Summary Responds with user id
// @Description Responds with user id
// @Tags auth
// @Produce json
// @Success 201 {object} types.AuthUserResponse
// @Failure 400 {object} httputil.HttpError
// @Failure 500 {object} httputil.HttpError
// @Router /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
}
