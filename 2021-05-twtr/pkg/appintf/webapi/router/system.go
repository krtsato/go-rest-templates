package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (f *facadeImpl) systemRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/health", f.system.Health)

	return r
}
