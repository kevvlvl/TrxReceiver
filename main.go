package main

import (
	"TrxReceiver/env"
	"TrxReceiver/rdb"
	"TrxReceiver/route"
	"github.com/rs/zerolog"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	redisClient := rdb.Instance()

	chiRouter := route.Router(&redisClient)
	chiRouter.ListenAndServe(env.ApiPort())
}
