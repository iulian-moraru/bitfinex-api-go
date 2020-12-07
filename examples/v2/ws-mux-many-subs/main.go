package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/mux"
)

func main() {
	srv := mux.
		New().
		AddClient()

	pairs := []string{}
	dat, err := ioutil.ReadFile("./testpairs.json")
	if err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(dat, &pairs); err != nil {
		log.Panic(err)
	}

	for _, pair := range pairs {
		tradePld := map[string]string{
			"event":   "subscribe",
			"channel": "trades",
			"symbol":  "t" + pair,
		}

		tickPld := map[string]string{
			"event":   "subscribe",
			"channel": "ticker",
			"symbol":  "t" + pair,
		}

		candlesPld := map[string]string{
			"event":   "subscribe",
			"channel": "candles",
			"key":     "trade:1m:t" + pair,
		}

		srv.Subscribe(tradePld)
		srv.Subscribe(tickPld)
		srv.Subscribe(candlesPld)
	}

	log.Fatal(srv.Listen(func(res []byte, err error) {
		if err != nil {
			log.Printf("error received: %s\n", err)
		}
		log.Printf("msg: %s\n", res)
	}))
}
