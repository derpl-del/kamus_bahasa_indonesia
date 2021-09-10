package server

import (
	"kamus_api/words"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graphql-go/handler"
)

func NewRoute(r *chi.Mux) *chi.Mux {
	graphQL := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	r.Use(middleware.Logger)
	r.Handle("/query", graphQL)
	r.Get("/api/arti/{kata}", words.RestApiGetWords)
	return r
}
