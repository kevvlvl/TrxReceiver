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
		port         = os.Getenv("API_PORT")
		redisHost    = os.Getenv("REDIS_HOST")
		redisPort, _ = strconv.ParseInt(os.Getenv("REDIS_PORT"), 10, 0)
		redisPass    = os.Getenv("REDIS_PASS")
	)

	redisClient := rdb.GetRedisClient(redisHost, int(redisPort), redisPass)

	chiRouter := route.Router(&redisClient)
	chiRouter.ListenAndServe(port)
}

func configureZeroLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
