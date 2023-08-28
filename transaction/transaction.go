package transaction

import (
	"TrxReceiver/rdb"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

const IntBase = 10

type Transaction struct {
	Id     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Value  float32 `json:"value"`
}

func GetTransaction(trxId int) (*Transaction, []byte) {

	redisClient := rdb.Client()

	trxStr := rdb.Get(intToStr(trxId), redisClient)
	trxByte := []byte(trxStr)

	var trx Transaction
	err := json.Unmarshal(trxByte, &trx)

	if err != nil {
		log.Error().Msgf("Error trying to unmarshall string byte to Transaction. %v", err)
	}

	log.Debug().Msgf("Found transaction %+v for ID %v", trx, trxId)

	return &trx, trxByte
}

func CreateTransaction(r *http.Request) {

	var s Transaction
	parseTransactionBody(r, &s)
	redisClient := rdb.Client()

	writeToRedis(s, redisClient)
	log.Debug().Msgf("Successfully created transaction %v", s.Id)
}

func UpdateTransaction(r *http.Request, trxId int) {

	var trxBody Transaction
	parseTransactionBody(r, &trxBody)

	if trxBody.Id != trxId {
		log.Error().Msg("The trxID in URL Path does not match request body Transaction ID")
	} else {

		redisClient := rdb.Client()
		trxDb := rdb.Get(intToStr(trxId), redisClient)
		log.Debug().Msgf("Found existing transaction %+v", trxDb)

		writeToRedis(trxBody, redisClient)
		log.Debug().Msgf("Successfully updated transaction %v", trxId)
	}
}

func parseTransactionBody(r *http.Request, s *Transaction) {

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		log.Error().Msgf("Error parsing Request Body: %s", err)
	}

	log.Debug().Msgf("Parsed JSON successfully: %+v", s)
}

func intToStr(i int) string {

	return strconv.FormatInt(int64(i), IntBase)
}

func writeToRedis(s Transaction, r *redis.Client) {
	jsonTrx, err := json.Marshal(s)

	if err != nil {
		log.Error().Msgf("Error trying to mashall Transaction to jsonTrx string: %v", err)
	}

	rdb.Set(intToStr(s.Id), string(jsonTrx[:]), r)
}
