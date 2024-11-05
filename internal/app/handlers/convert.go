package handlers

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/jvcosta-dev/go-currency-exchange/internal/app/validations"
	"github.com/jvcosta-dev/go-currency-exchange/internal/database"
)

type ConversionResponse struct {
	From       string  `json:"from"`
	To         string  `json:"to"`
	Amount     float64 `json:"amount"`
	Result     float64 `json:"result"`
	Rate       float64 `json:"rate"`
	LastUpdate string  `json:"last_update"`
}

func Convert(w http.ResponseWriter, r *http.Request) {
	from := r.PathValue("from")
	to := r.PathValue("to")
	amountstr := r.PathValue("amount")
	if from == "" || to == "" || amountstr == "" {
		http.Error(w, "Missing required parameters: 'from', 'to', and 'amount' must be provided", http.StatusBadRequest)
		return
	}

	if !validations.IsValidCurrency(from) || !validations.IsValidCurrency(to) {
		http.Error(w, "Invalid currency code", http.StatusBadRequest)
		return
	}

	amount, err := validations.ValidateAmount(amountstr)
	if err != nil {
		http.Error(w, "Invalid amount format", http.StatusBadRequest)
		return
	}

	if amount > 1e6 {
		http.Error(w, "Invalid amount size", http.StatusBadRequest)
		return
	}

	lastUpdate, err := database.GetLastUpdate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fromRate, err := database.GetRate(from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toRate, err := database.GetRate(to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := amount * (toRate / fromRate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ConversionResponse{
		From:       from,
		To:         to,
		Amount:     amount,
		Result:     math.Round(result*1000) / 1000,
		Rate:       fromRate,
		LastUpdate: lastUpdate,
	})
}
