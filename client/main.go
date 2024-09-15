package main

import (
	"client/src"
	"log"
	"time"
)

func main() {
	for {
		time.Sleep(10 * time.Second)
		log.Print(src.GetBid())
	}
}
