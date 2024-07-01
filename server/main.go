package main

import (
	"encoding/json"
	"log"
	"net/http"
	config "server/configs"
	"server/infra/db"
	"server/src"
)

var port = ":" + *config.GetPort()

func main() {
	http.HandleFunc("/cotacao", quotationHandler)
	db.CreateDB()
	db.InsertQuotation(&src.Quotation{Bid: "21.4"})
	http.ListenAndServe(port, nil)
}

func quotationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	w.Write(callApi())
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
