package route

import (
	"TrxReceiver/rdb"
	"TrxReceiver/transaction"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"strconv"
)

type ChiRouter struct {
	Router *chi.Mux
	Trx    *transaction.Trx
}

func Router(redisClient *rdb.RedisDB) ChiRouter {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	return ChiRouter{
		Router: r,
		Trx: &transaction.Trx{
			Redis: redisClient,
		},
	}
}

func (router *ChiRouter) ListenAndServe() {

	router.handleRoutes()

	port := os.Getenv("API_PORT")

	if port == "" {
		port = "3000"
	}

	log.Info().Msgf("Port: %v", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.Router)

	if err != nil {
		return
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

				trxId := parseTrxId(r)
				log.Info().Msgf("GET on Stock ID %v", trxId)

				_, trxBytes := router.Trx.GetTransaction(trxId)

				_, err := w.Write(trxBytes)
				if err != nil {
					return
				}
			})

			r.Put("/", func(w http.ResponseWriter, r *http.Request) {

				stockId := parseTrxId(r)
				log.Info().Msgf("PUT on Stock ID %v", stockId)

				router.Trx.UpdateTransaction(r, stockId)
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
