package transaction

import (
	"TrxReceiver/rdb"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

const IntBase = 10

type Trx struct {
	Redis *rdb.RedisDB
}

type Stock struct {
	Id     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Value  float32 `json:"value"`
}

func (t *Trx) GetTransaction(trxId int) (*Stock, []byte) {

	trxStr := t.Redis.Get(intToStr(trxId))
	trxByte := []byte(trxStr)

	var trx Stock
	err := json.Unmarshal(trxByte, &trx)

	if err != nil {
		log.Error().Msgf("Error trying to unmarshall string byte to Stock. %v", err)
	}

	log.Debug().Msgf("Found transaction %+v for ID %v", trx, trxId)

	return &trx, trxByte
}

func (t *Trx) CreateTransaction(r *http.Request) {

	var s Stock
	parseTransactionBody(r, &s)

	t.writeToRedis(s)
	log.Debug().Msgf("Successfully created transaction %v", s.Id)
}

func (t *Trx) UpdateTransaction(r *http.Request, trxId int) {

	var trxBody Stock
	parseTransactionBody(r, &trxBody)

	if trxBody.Id != trxId {
		log.Error().Msg("The trxID in URL Path does not match request body Stock ID")
	} else {

		trxDb := t.Redis.Get(intToStr(trxId))
		log.Debug().Msgf("Found existing transaction %+v", trxDb)

		t.writeToRedis(trxBody)
		log.Debug().Msgf("Successfully updated transaction %v", trxId)
	}
}

func (t *Trx) writeToRedis(s Stock) {
	jsonTrx, err := json.Marshal(s)

	if err != nil {
		log.Error().Msgf("Error trying to mashall Stock to jsonTrx string: %v", err)
	}

	t.Redis.Set(intToStr(s.Id), string(jsonTrx[:]), 0)
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
