package main

import (
	"log"
	//"time"

	"github.com/joho/godotenv"
	"github.com/jvcosta-dev/go-currency-exchange/internal/app"
	"github.com/jvcosta-dev/go-currency-exchange/internal/database"
)

func main() {
	err := godotenv.Load("internal/config/.env.local")
	if err != nil {
		log.Fatal(err)
	}
	database.Init()
	//go database.StartUpdatingRates(6 * time.Hour)
	app.Server()
}
