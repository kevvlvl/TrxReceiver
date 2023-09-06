package route

import (
	"TrxReceiver/rdb"
	"TrxReceiver/transaction"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	metrics "github.com/m8as/go-chi-metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Router(redisClient *rdb.RedisDB) ChiRouter {

	r := chi.NewRouter()

	r.Use(middleware.CleanPath)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)
	r.Use(metrics.SetRequestDuration)
	r.Use(metrics.IncRequestCount)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	return ChiRouter{
		Router: r,
		Trx: &transaction.Trx{
			Redis: redisClient,
		},
	}
}

func (router *ChiRouter) ListenAndServe(port string) {

	router.handleRoutes()

	if port == "" {
		port = "3000"
	}

	log.Info().Msgf("Port: %v", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.Router); err != nil {
		log.Info().Msgf("Listen and serve error: %v", err)
	}
}

func (router *ChiRouter) handleRoutes() {

	router.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("Root path"))
		if err != nil {
			return
		}
	})

	router.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("Up and Healthy"))
		if err != nil {
			return
		}
	})

	router.Router.Route("/trx", func(r chi.Router) {

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {

			router.Trx.CreateTransaction(r)
		})

		r.Route("/{trxID}", func(r chi.Router) {

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {

				trxId := parseStockId(r)
				log.Info().Msgf("GET on Stock ID %v", trxId)

				_, trxBytes := router.Trx.GetTransaction(trxId)

				_, err := w.Write(trxBytes)
				if err != nil {
					return
				}
			})

			r.Put("/", func(w http.ResponseWriter, r *http.Request) {

				stockId := parseStockId(r)
				log.Info().Msgf("PUT on Stock ID %v", stockId)

				router.Trx.UpdateTransaction(r, stockId)
			})
		})
	})

	router.Router.Handle("/prometheus", promhttp.Handler())
}

func parseStockId(r *http.Request) string {
	return chi.URLParam(r, "trxID")
}
