package stock

import (
	"TrxReceiver/rdb"
	"encoding/json"
	"github.com/redis/go-redis/v9"
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

	redisClient := rdb.Client()

	stockStr := rdb.Get(intToStr(stockId), redisClient)
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
	redisClient := rdb.Client()

	writeToRedis(s, redisClient)
	log.Debug().Msgf("Successfully created stock %s", s.Id)
}

func UpdateTransaction(r *http.Request, stockId int) {

	var stockBody Stock
	parseTransactionBody(r, &stockBody)

	if stockBody.Id != stockId {
		log.Error().Msg("The stockId in URL Path does not match request body Stock ID")
	} else {

		redisClient := rdb.Client()
		stockDb := rdb.Get(intToStr(stockId), redisClient)
		log.Debug().Msgf("Found existing stock %+v", stockDb)

		writeToRedis(stockBody, redisClient)
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

func writeToRedis(s Stock, r *redis.Client) {
	jsonStock, err := json.Marshal(s)

	if err != nil {
		log.Error().Msgf("Error trying to mashall Stock to jsonStock string: %v", err)
	}

	rdb.Set(intToStr(s.Id), string(jsonStock[:]), r)
}
