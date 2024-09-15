package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	config "server/configs"
	"server/infra/db"
	"server/src"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var port = ":" + *config.GetPort()

var (
	ErrCallingExternalApi = errors.New("chamada falhou, tente novamente")
	ErrUnmarshal          = errors.New("erro durante decodificação do JSON")
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/cotacao", quotationHandler)

	db.CreateDB()
	http.ListenAndServe(port, router)
}

func quotationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, err := callExternalApi()
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func callExternalApi() ([]byte, error) {
	res, err := src.GetQuotation()
	if err != nil {
		return nil, ErrCallingExternalApi
	}

	q := &src.Quotation{Bid: res.USDBRL.Bid}
	db.InsertQuotation(&src.Quotation{Bid: q.Bid})

	txt, err := json.Marshal(q)
	if err != nil {
		return nil, ErrUnmarshal
	}

	return txt, nil
}
