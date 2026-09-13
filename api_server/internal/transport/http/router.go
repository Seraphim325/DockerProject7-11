package http

import (
	"api_server/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/values/all", func(r chi.Router) {
		r.Post("/", h.IndexHandler.SaveIndex)
		r.Get("/", h.IndexHandler.GetAllIndexes)
	})

	r.Route("/api/values/current", func(r chi.Router) {
		r.Post("/", h.ValueHandler.SaveValue)
		r.Get("/", h.ValueHandler.GetAllValues)
	})

	return r
}
