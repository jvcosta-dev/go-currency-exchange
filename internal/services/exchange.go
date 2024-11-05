package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type LatestRatesResponse struct {
	ConversionRates   map[string]float64 `json:"conversion_rates"`
	TimeLastUpdateUtc string             `json:"time_last_update_utc"`
}

func GetLatestRates() (*LatestRatesResponse, error) {
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/USD", apiKey)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("%s", body)
		return nil, fmt.Errorf("failed to get exchange rate: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var latestRatesResponse LatestRatesResponse
	if err := json.Unmarshal(body, &latestRatesResponse); err != nil {
		return nil, err
	}

	return &latestRatesResponse, nil
}
