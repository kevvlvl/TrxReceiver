package stock

import (
	"TrxReceiver/rdb"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

const INT_BASE = 10

type Stock struct {
	Id     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Value  float32 `json:"value"`
}

func GetTransaction(r *http.Request, stockId int) {

	redis := rdb.Client()

	s := rdb.Get(strconv.FormatInt(int64(stockId), INT_BASE), redis)

	log.Debug().Msgf("Found stock %+v for ID %v", s, stockId)
}

func CreateTransaction(r *http.Request) {

	var s Stock
	parseTransactionBody(r, &s)
	redis := rdb.Client()

	json, err := json.Marshal(s)

	if err != nil {
		log.Error().Msgf("Error trying to mashall Stock to json string: %v", err)
	}

	rdb.Set(strconv.FormatInt(int64(s.Id), INT_BASE), string(json[:]), redis)
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
