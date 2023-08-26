package http

import (
	"TrxReceiver/stock"
	"github.com/go-chi/chi/v5"
	"net/http"
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

		r.Post("/", stock.CreateTransaction)
		r.Route("/{stockID}", func(r chi.Router) {
			r.Use(stock.TransactionCtx)
			r.Put("/", stock.UpdateTransaction)
		})
	})
}
