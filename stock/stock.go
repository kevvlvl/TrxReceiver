package stock

import (
	"TrxReceiver/rdb"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

const IntBase = 10

type Stock struct {
	Id     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Value  float32 `json:"value"`
}

func GetTransaction(stockId int) (*Stock, []byte) {

	redis := rdb.Client()

	stockStr := rdb.Get(intToStr(stockId), redis)
	stockByte := []byte(stockStr)

	var stock Stock
	err := json.Unmarshal(stockByte, &stock)

	if err != nil {
		log.Error().Msgf("Error trying to unmarshall string byte to Stock. %v", err)
	}

	log.Debug().Msgf("Found stock %+v for ID %v", stock, stockId)

	return &stock, stockByte
}

func CreateTransaction(r *http.Request) {

	var s Stock
	parseTransactionBody(r, &s)
	redis := rdb.Client()

	jsonStock, err := json.Marshal(s)

	if err != nil {
		log.Error().Msgf("Error trying to mashall Stock to jsonStock string: %v", err)
	}

	rdb.Set(intToStr(s.Id), string(jsonStock[:]), redis)

	log.Debug().Msgf("Successfully created stock %s", s.Id)
}

func UpdateTransaction(r *http.Request, stockId int) {

	var stockBody Stock
	parseTransactionBody(r, &stockBody)

	if stockBody.Id != stockId {
		log.Error().Msg("The stockId in URL Path does not match request body Stock ID")
	} else {

		redis := rdb.Client()
		stockDb := rdb.Get(intToStr(stockId), redis)
		log.Debug().Msgf("Found existing stock %+v", stockDb)

		jsonStock, err := json.Marshal(stockBody)

		if err != nil {
			log.Error().Msgf("Error trying to mashall Stock to jsonStock string: %v", err)
		}

		rdb.Set(intToStr(stockId), string(jsonStock[:]), redis)

		log.Debug().Msgf("Successfully updated stock %s", stockId)
	}
}

func parseTransactionBody(r *http.Request, s *Stock) {

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		log.Error().Msgf("Error parsing Request Body: %s", err)
	}

	log.Debug().Msgf("Parsed JSON successfully: %+v", s)
}

func intToStr(i int) string {

	return strconv.FormatInt(int64(i), IntBase)
}
