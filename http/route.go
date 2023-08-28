package http

import (
	"TrxReceiver/transaction"
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

			transaction.CreateTransaction(r)
		})

		r.Route("/{trxID}", func(r chi.Router) {

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {

				trxId := parseTrxId(r)
				log.Info().Msgf("GET on Transaction ID %v", trxId)

				_, trxBytes := transaction.GetTransaction(trxId)

				_, err := w.Write(trxBytes)
				if err != nil {
					return
				}
			})

			r.Put("/", func(w http.ResponseWriter, r *http.Request) {

				stockId := parseTrxId(r)
				log.Info().Msgf("PUT on Transaction ID %v", stockId)

				transaction.UpdateTransaction(r, stockId)
			})
		})
	})
}

func parseTrxId(r *http.Request) int {
	stockId, err := strconv.ParseInt(chi.URLParam(r, "trxID"), 10, 32)

	if err != nil {
		log.Error().Msgf("Error parsing the URL param trxID: %s", err)
	}

	return int(stockId)
}
