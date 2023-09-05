package main

import (
	"TrxReceiver/rdb"
	"TrxReceiver/route"
	"github.com/rs/zerolog"
	"os"
)

func main() {

	configureZeroLog()

	redisClient := rdb.Instance()

	chiRouter := route.Router(&redisClient)
	chiRouter.ListenAndServe(os.Getenv("API_PORT"))
}

func configureZeroLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
