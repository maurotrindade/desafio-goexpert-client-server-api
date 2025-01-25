package main

import (
	"client/src"
	"log"
	"time"
)

func main() {
	for {
		time.Sleep(10 * time.Second)

		bid, err := src.GetBid()
		if err != nil {
			log.Print(err)
		}
		if bid != nil {
			src.CreateFile(bid.Bid)
		}
	}
}
