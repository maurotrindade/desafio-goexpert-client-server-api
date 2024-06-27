package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	config "server/configs"
	"time"
)

func main() {
	log.Print(*config.GetQuotationAddress())
	res, _ := getQuotation()
	log.Print(res)
}

type QuotationRes struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func getQuotation() (*QuotationRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	log.Printf("Iniciando chamada no endereço: %s", *config.GetQuotationAddress())
	req, err := http.NewRequestWithContext(ctx, "GET", *config.GetQuotationAddress(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: verificar problema com TLS no container
	if *config.GetEnv() == "DEV" {
		cfg := &tls.Config{
			InsecureSkipVerify: true,
		}
		http.DefaultClient.Transport = &http.Transport{
			TLSClientConfig: cfg,
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erro ao fazer a chamada: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Erro ao ler resposta: %v", err)
		return nil, err
	}

	var quotation QuotationRes
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		log.Printf("Erro ao parsear json: %v", err)
		return nil, err
	}

	return &quotation, nil
}
