package router

import "github.com/go-chi/chi/v5"

func New() (chi.Router, error) {
	router := chi.NewRouter()
	return router, nil
}
