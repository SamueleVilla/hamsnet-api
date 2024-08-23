package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/samuelevilla/hasnet-api/internal/httputil"
	"github.com/samuelevilla/hasnet-api/internal/middleware"
	"github.com/samuelevilla/hasnet-api/internal/store"
)

type HamsterHandler struct {
	store store.Store
}

func NewHamsterHandler(store store.Store) *HamsterHandler {
	return &HamsterHandler{
		store: store,
	}
}

func (h *HamsterHandler) RegisterRoutes(router chi.Router) {
	// hamsters
	router.Route("/hamsters", func(r chi.Router) {
		r.Get("/feed", h.HandleHamstersFeed)
		r.Get("/{id}", h.HandleHamsterById)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth)
			r.Post("/", h.HandleCreateHamsterPost)
		})
	})
}

// HandleHamstersFeed godoc
// @Summary Responds with a list of hamster posts
// @Description Responds with a list of hamster posts
// @Tags feed
// @Produce json
// @Success 200 {object} store.HamsterPost
// @Failure 500 {object} httputil.HttpError
// @Router /hamsters/feed [get]
func (h *HamsterHandler) HandleHamstersFeed(w http.ResponseWriter, r *http.Request) {
	data, err := h.store.FindHamstersFeed(r.Context())
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, err.Error())
	}

	httputil.WriteJSON(w, http.StatusOK, data)
}

// HandlerHamsterById godoc
// @Summary Responds with the hamster post with the given id
// @Description Responds with the hamster post with the given id
// @Tags feed
// @Produce json
// @Success 200 {object} store.HamsterPost
// @Failure 500 {object} httputil.HttpError
// @Router /hamsters/{id} [get]
func (h *HamsterHandler) HandleHamsterById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := h.store.FindHamsterById(r.Context(), id)
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, err.Error())
	}

	httputil.WriteJSON(w, http.StatusOK, data)
}

// HandlerCreateHamsterPost godoc
// @Summary Responds with created hamster post
// @Description Responds with created hamster post
// @Tags feed
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} httputil.HttpError
// @Failure 400 {object} httputil.HttpError
// @Failure 500 {object} httputil.HttpError
// @Router /hamsters [post]
func (h *HamsterHandler) HandleCreateHamsterPost(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("userId").(string)
	content := r.FormValue("content")
	if content == "" {
		httputil.WriteError(w, http.StatusBadRequest, "missing content")
	}

	fmt.Printf("userId from handler: %s\n", userId)
	post := &store.CreateHamsterPost{
		AuthorId: userId,
		Content:  content,
	}

	err := h.store.CreateHamsterPost(r.Context(), post)
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, err.Error())
	}

	data := map[string]string{"message": "hamster post created"}
	httputil.WriteJSON(w, http.StatusCreated, data)
}
