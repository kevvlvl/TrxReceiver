package transaction

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (t *Trx) GetAll() []byte {

	keys, err := t.Redis.Client.Keys(t.Redis.Context, "*").Result()

	if err != nil {
		log.Error().Msgf("Failed to obtain all keys. %v", err)
		return nil
	}

	var stocks []Stock
	for i := 0; i < len(keys); i++ {

		currentStr := t.Redis.Get(keys[i])
		currentStock, _ := unmarshalStock(currentStr)

		stocks = append(stocks, currentStock)
	}

	log.Info().Msgf("Found %v number of entries in Redis", len(stocks))

	return nil
}

func (t *Trx) GetTransaction(stockId string) []byte {

	stockStr := t.Redis.Get(stockId)
	_, b := unmarshalStock(stockStr)

	return b
}

func (t *Trx) CreateTransaction(r *http.Request) {

	var s Stock
	parseTransactionBody(r, &s)

	t.writeToRedis(&s)
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

		t.writeToRedis(&stockBody)
		log.Debug().Msgf("Successfully updated stock %v", stockId)
	}
}

func (t *Trx) writeToRedis(f FinancialInstrument) {
	t.Redis.Set(f.IdStr(), string(f.AsBytes()[:]), 0)
}

func parseTransactionBody(r *http.Request, s *Stock) {

	err := json.NewDecoder(r.Body).Decode(s)

	if err != nil {
		log.Error().Msgf("Error parsing Request Body: %s", err)
	}

	log.Debug().Msgf("Parsed JSON successfully: %+v", s)
}

func unmarshalStock(str string) (Stock, []byte) {

	b := []byte(str)
	var s Stock

	err := json.Unmarshal(b, &s)

	if err != nil {
		log.Error().Msgf("Error trying to unmarshall string byte to Stock. %v", err)
	}

	log.Debug().Msgf("Found stock %+v for ID %v", s, s.Id)

	return s, b
}
