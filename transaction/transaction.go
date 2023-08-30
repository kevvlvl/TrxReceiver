package transaction

import (
	"TrxReceiver/rdb"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Trx struct {
	Redis *rdb.RedisDB
}

func (t *Trx) GetTransaction(stockId string) (*Stock, []byte) {

	trxStr := t.Redis.Get(stockId)
	trxByte := []byte(trxStr)

	var trx Stock
	err := json.Unmarshal(trxByte, &trx)

	if err != nil {
		log.Error().Msgf("Error trying to unmarshall string byte to Stock. %v", err)
	}

	log.Debug().Msgf("Found stock %+v for ID %v", trx, stockId)

	return &trx, trxByte
}

func (t *Trx) CreateTransaction(r *http.Request) {

	var s Stock
	parseTransactionBody(r, &s)

	t.writeToRedis(s)
	log.Debug().Msgf("Successfully created transaction %v", s.Id)
}

func (t *Trx) UpdateTransaction(r *http.Request, stockId string) {

	var stockBody Stock
	parseTransactionBody(r, &stockBody)

	if stockBody.IdStr() != stockId {
		log.Error().Msg("The trxID in URL Path does not match request body Stock ID")
	} else {

		stockDb := t.Redis.Get(stockBody.IdStr())
		log.Debug().Msgf("Found existing stock %+v", stockDb)

		t.writeToRedis(stockBody)
		log.Debug().Msgf("Successfully updated stock %v", stockId)
	}
}

func (t *Trx) writeToRedis(s Stock) {
	t.Redis.Set(s.IdStr(), string(s.AsBytes()[:]), 0)
}

func parseTransactionBody(r *http.Request, s *Stock) {

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		log.Error().Msgf("Error parsing Request Body: %s", err)
	}

	log.Debug().Msgf("Parsed JSON successfully: %+v", s)
}
