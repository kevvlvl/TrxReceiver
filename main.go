package main

import (
	"TrxReceiver/http"
	"github.com/rs/zerolog"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	r := http.Router()
	http.HandleRoutes(r)
	http.Serve(r)
}
