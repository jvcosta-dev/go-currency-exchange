package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jvcosta-dev/go-currency-exchange/internal/app/validations"
	"github.com/jvcosta-dev/go-currency-exchange/internal/database"
)

type RateResponse struct {
	Base       string  `json:"base"`
	Rate       float64 `json:"rate"`
	LastUpdate string  `json:"last_update"`
}

func Rate(w http.ResponseWriter, r *http.Request) {
	base := r.PathValue("base")
	if base == "" {
		http.Error(w, "Missing required parameters: 'base' must be provided", http.StatusBadRequest)
		return
	}

	if !validations.IsValidCurrency(base) {
		http.Error(w, "Invalid currency code", http.StatusBadRequest)
		return
	}

	rate, err := database.GetRate(base)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastUpdate, err := database.GetLastUpdate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RateResponse{
		Base:       base,
		Rate:       rate,
		LastUpdate: lastUpdate,
	})
}

type RatesResponse struct {
	Rates      map[string]float64 `json:"rates"`
	LastUpdate string             `json:"last_update"`
}

func Rates(w http.ResponseWriter, r *http.Request) {
	rates, err := database.GetRates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastUpdate, err := database.GetLastUpdate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RatesResponse{
		Rates:      rates,
		LastUpdate: lastUpdate,
	})

}
