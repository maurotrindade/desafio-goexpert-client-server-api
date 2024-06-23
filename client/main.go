package main

import (
	config "client/configs"
	"log"
)

func main() {
	log.Print(*config.GetPort())
}
