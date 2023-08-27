package stock

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Stock struct {
	Id     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Value  float32 `json:"value"`
}

func CreateTransaction(r *http.Request) {

	var s Stock
	parseTransactionBody(r, &s)
}

func UpdateTransaction(r *http.Request, stockId int) {

	var s Stock
	parseTransactionBody(r, &s)
}

func parseTransactionBody(r *http.Request, s *Stock) {

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		log.Error().Msgf("Error parsing Request Body: %s", err)
	}

	log.Debug().Msgf("Parsed JSON successfully: %+v", s)
}
