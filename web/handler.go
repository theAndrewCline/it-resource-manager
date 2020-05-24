package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/theAndrewCline/it-resource-manager/postgres"
	"github.com/theAndrewCline/it-resource-manager/types"

	"github.com/go-chi/chi"
)

// NewHandler Constructore for handler
func NewHandler(store *postgres.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)
	h.Route("/Owners", func(r chi.Router) {
		r.Get("/", h.OwnerList())
	})

	return h
}

// Handler for http requests
type Handler struct {
	*chi.Mux
	store types.Store
}

// OwnerList handler for getting all the owners
func (h *Handler) OwnerList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cc, err := h.store.Computers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(cc)
	}
}
