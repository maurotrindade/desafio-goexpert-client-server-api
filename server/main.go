package main

import (
	"encoding/json"
	"log"
	"net/http"
	config "server/configs"
	"server/src"
)

var port = ":" + *config.GetPort()

func main() {
	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		w.Write(callApi())
	})
	http.ListenAndServe(port, nil)
}

func callApi() []byte {
	res, err := src.GetQuotation()
	if err != nil {
		log.Fatal("Chamada falhou, tente novamente")
	}

	q := &src.Quotation{Bid: res.USDBRL.Bid}
	txt, err := json.Marshal(q)
	if err != nil {
		log.Fatal("Deu ruim na Patrulha Canina")
	}

	return txt
}
