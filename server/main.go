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
	errCallingExternalApi = errors.New("chamada falhou, tente novamente")
	errUnmarshal          = errors.New("erro durante decodificação do JSON")
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
		// return
	}

	json, err := callExternalApi()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(json)
}

func callExternalApi() ([]byte, error) {
	res, err := src.GetQuotation()
	if err != nil {
		return nil, errCallingExternalApi
	}

	q := &src.Quotation{Bid: res.USDBRL.Bid}

	err = db.InsertQuotation(&src.Quotation{Bid: q.Bid})
	if err != nil {
		return nil, errCallingExternalApi
	}

	txt, err := json.Marshal(q)
	if err != nil {
		return nil, errUnmarshal
	}

	return txt, nil
}
