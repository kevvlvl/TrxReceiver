package http

import (
	"TrxReceiver/stock"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func HandleRoutes(router *chi.Mux) {

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("Root path"))
		if err != nil {
			return
		}
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("Up and Healthy"))
		if err != nil {
			return
		}
	})

	router.Route("/trx", func(r chi.Router) {

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {

			stock.CreateTransaction(r)
		})

		r.Put("/{stockID}", func(w http.ResponseWriter, r *http.Request) {

			stockId, err := strconv.ParseInt(chi.URLParam(r, "stockID"), 10, 32)

			if err != nil {
				log.Error().Msgf("Error parsing the URL param stockID: %s", err)
			}

			log.Info().Msgf("PUT on Stock ID %v", stockId)
			stock.UpdateTransaction(r, int(stockId))
		})
	})
}
