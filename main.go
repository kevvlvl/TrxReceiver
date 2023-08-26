package main

import "TrxReceiver/http"

func main() {

	r := http.Router()
	http.HandleRoutes(r)
	http.Serve(r)
}
