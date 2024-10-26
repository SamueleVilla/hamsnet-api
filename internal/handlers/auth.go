package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samuelevilla/hasnet-api/internal/httputil"
	"github.com/samuelevilla/hasnet-api/internal/store"
	"github.com/samuelevilla/hasnet-api/internal/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	store  store.Store
	secret string
}

func NewAuthHandler(store store.Store, secret string) *AuthHandler {
	return &AuthHandler{
		store:  store,
		secret: secret,
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
// @Failure 401 {object} httputil.HttpError
// @Failure 500 {object} httputil.HttpError
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	usernameOrEmail := r.FormValue("usernameOrEmail")
	password := r.FormValue("password")

	// find user by username or email
	user, err := h.store.FindUserByUsernameOrEmail(r.Context(), usernameOrEmail)
	if err != nil {
		httputil.WriteError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		httputil.WriteError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// Create the Claims
	roles := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = role.RoleName
	}
	// Set the expiration time for the token 24 hours
	// TODO: Move this to a config file
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := types.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "hasnet-api",
			Subject:   user.Id,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		UserId:   user.Id,
		Username: user.Username,
		Email:    user.Email,
		Roles:    roles,
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(h.secret))

	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "error generating token")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, types.AuthUserResponse{UserId: user.Id, Token: tokenString})
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
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if username == "" || email == "" || password == "" {
		httputil.WriteError(w, http.StatusBadRequest, "missing required fields")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "error hashing password")
		return
	}

	createUser := &store.CreateUser{
		Username:       username,
		Email:          email,
		HashedPassword: string(hashedPassword),
	}

	userId, err := h.store.CreateUser(r.Context(), createUser)
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "error creating user")
		return
	}

	// Set the expiration time for the token 24 hours
	// TODO: Move this to a config file
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := types.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "hasnet-api",
			Subject:   userId,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		UserId:   userId,
		Username: username,
		Email:    email,
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(h.secret))

	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "error generating token")
		return
	}

	httputil.WriteJSON(w, http.StatusCreated, types.AuthUserResponse{UserId: userId, Token: tokenString})
}
