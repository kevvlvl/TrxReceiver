package main

import (
	"TrxReceiver/rdb"
	"TrxReceiver/route"
	"github.com/rs/zerolog"
	"os"
	"strconv"
)

func main() {

	configureZeroLog()

	var (
		redisHost    = os.Getenv("REDIS_HOST")
		redisPort, _ = strconv.ParseInt(os.Getenv("REDIS_PORT"), 10, 0)
		redisPass    = os.Getenv("REDIS_PASS")
	)

	redisClient := rdb.GetRedisClient(redisHost, int(redisPort), redisPass)

	chiRouter := route.Router(&redisClient)
	chiRouter.ListenAndServe()
}

func configureZeroLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
