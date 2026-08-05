package converter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type exchangeResponse struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

func Convert(amount float64, from, to string) (float64, error) {
	url := fmt.Sprintf("https://api.frankfurter.dev/v1/latest?base=%s&symbols=%s", from, to)

	response, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("erro ao consultar a API: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API retornou status %d (verifique os códigos de moeda)", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	var data exchangeResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	rate, ok := data.Rates[to]
	if !ok {
		return 0, fmt.Errorf("moeda de destino '%s' não encontrada na resposta", to)
	}

	return amount * rate, nil
}