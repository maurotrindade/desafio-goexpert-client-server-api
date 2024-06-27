package main

import (
	config "client/configs"
	"log"
	"time"
)

func main() {
	for {
		time.Sleep(10 * time.Second)
		log.Print(*config.GetPort())
	}
}
