package main

import (
	"client/src"
	"log"
	"time"
)

func main() {
	for {
		time.Sleep(10 * time.Second)

		err, bid := src.GetBid()
		if err != nil {
			log.Print(err)
		}
		log.Print(bid)
	}
}
