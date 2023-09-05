package route

import (
	"TrxReceiver/transaction"
	"github.com/go-chi/chi/v5"
)

type ChiRouter struct {
	Router *chi.Mux
	Trx    *transaction.Trx
}
