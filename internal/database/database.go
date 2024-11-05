package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/jvcosta-dev/go-currency-exchange/internal/services"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "./main.db")

	if err != nil {
		log.Fatal(err)
	}

	Migrate()
}

func UpdateRates(rates map[string]float64, lastUpdate string) error {
	// Inicia uma transação para garantir consistência ao atualizar taxas e metadata
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Atualiza cada taxa individualmente
	for currency, rate := range rates {
		_, err := tx.Exec(`
            INSERT INTO rates (currency, rate) 
            VALUES (?, ?) 
            ON CONFLICT(currency) DO UPDATE SET rate = excluded.rate`, currency, rate)
		if err != nil {
			return err
		}
	}

	// Atualiza a data de última atualização na tabela metadata
	_, err = tx.Exec(`
        INSERT INTO metadata (key, value) 
        VALUES ('last_update', ?) 
        ON CONFLICT(key) DO UPDATE SET value = excluded.value`, lastUpdate)
	if err != nil {
		return err
	}

	// Confirma a transação
	return tx.Commit()
}

func GetServiceRates() {
	ratesResponse, err := services.GetLatestRates()
	if err != nil {
		log.Printf("Error fetching latest rates: %v", err)
		return
	}

	err = UpdateRates(ratesResponse.ConversionRates, ratesResponse.TimeLastUpdateUtc)
	if err != nil {
		log.Printf("Error updating exchange rates in the database: %v", err)
	}
}

func StartUpdatingRates(interval time.Duration) {
	GetServiceRates()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		GetServiceRates()
	}
}

// Retorna a última data de atualização global das taxas
func GetLastUpdate() (string, error) {
	var lastUpdate string
	query := "SELECT value FROM metadata WHERE key = 'last_update'"
	err := DB.QueryRow(query).Scan(&lastUpdate)
	return lastUpdate, err
}

// Retorna todas as taxas de câmbio
func GetRates() (map[string]float64, error) {
	rates := make(map[string]float64)

	query := "SELECT currency, rate FROM rates"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currency string
		var rate float64
		if err := rows.Scan(&currency, &rate); err != nil {
			return nil, err
		}
		rates[currency] = rate
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rates, nil
}

// Retorna a taxa para uma moeda específica
func GetRate(currency string) (float64, error) {
	var rate float64
	query := "SELECT rate FROM rates WHERE currency = ?"
	err := DB.QueryRow(query, currency).Scan(&rate)
	return rate, err
}
