package src

import (
	config "client/configs"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

type Quotation struct {
	Bid string `json:"bid"`
}

var errTimeout = errors.New("tempo de resposta da API excedido")

func GetBid() (*Quotation, error) {
	ctx, cancel := context.WithTimeoutCause(context.Background(), 300*time.Millisecond, errTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", *config.GetServerAddress(), nil)
	if err != nil {
		log.Printf("erro ao criar contexto: %v", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
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
