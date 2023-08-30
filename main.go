package main

import (
	"TrxReceiver/rdb"
	"TrxReceiver/route"
	"github.com/rs/zerolog"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	redisClient := rdb.GetRedisClient("localhost", 6379, "")

	chiRouter := route.Router(&redisClient)
	chiRouter.ListenAndServe()
}
