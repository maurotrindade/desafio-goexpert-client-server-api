package src

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	config "server/configs"
	"time"
)

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

var errTimeout = errors.New("tempo de resposta do servidor excedido")

func GetQuotation() (*QuotationRes, error) {
	ctx, cancel := context.WithTimeoutCause(context.Background(), 4000*time.Millisecond, errTimeout)
	defer cancel()

	log.Printf("Chamada recebida no endereço: %s\n", *config.GetQuotationAddress())
	req, err := http.NewRequestWithContext(ctx, "GET", *config.GetQuotationAddress(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: verificar problema com TLS no container (pode ser versão do Go)
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
		log.Print(err)
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
