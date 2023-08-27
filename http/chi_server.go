package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func Router() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	return r
}

func Serve(router *chi.Mux) {

	port := os.Getenv("API_PORT")

	if port == "" {
		port = "3000"
	}

	log.Info().Msgf("Port: %v", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)

	if err != nil {
		return
	}
}
