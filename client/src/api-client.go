package src

import (
	config "client/configs"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Quotation struct {
	Bid string `json:"bid"`
}

func GetBid() (*Quotation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", *config.GetServerAddress(), nil)
	if err != nil {
		log.Printf("erro ao criar contexto: %v", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("erro ao fazer a chamada ao servidor: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("erro ao ler resposta: %v", err)
		return nil, err
	}

	var quotation Quotation
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		log.Printf("erro ao parsear json: %v", err)
		return nil, err
	}

	return &quotation, nil
}
