package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/samuelevilla/hasnet-api/internal/types"
)

type HamsterHandler struct {
}

func NewHamsterHandler() *HamsterHandler {
	return &HamsterHandler{}
}

func (h *HamsterHandler) RegisterRoutes(router *chi.Mux) {
	// hamsters
	router.Get("/hamsters", h.HandleGetHamsters) // GET http://localhost:8080/hamsters
}

func (h *HamsterHandler) HandleGetHamsters(w http.ResponseWriter, r *http.Request) {

	hamsters := []types.Hamster{
		{
			Id:          uint64(rand.Intn(100000)),
			ImgUrl:      "https://images.unsplash.com/photo-1719937206094-8de79c912f40?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			Content:     "This is a hamster",
			Visibility:  "public",
			LikesCount:  100,
			SharesCount: 50,
			Comments: []string{
				"Nice",
				"Wow",
			},
			AuthorID:  1,
			CreatedAt: time.Now().String(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hamsters)
}
