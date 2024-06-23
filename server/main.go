package main

import (
	"log"
	config "server/configs"
)

func main() {
	log.Print(*config.GetPort())
}
