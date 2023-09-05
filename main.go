package main

import (
	"TrxReceiver/rdb"
	"TrxReceiver/route"
	"github.com/rs/zerolog"
	"os"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	redisClient := rdb.Instance()

	chiRouter := route.Router(&redisClient)
	chiRouter.ListenAndServe(os.Getenv("API_PORT"))
}
