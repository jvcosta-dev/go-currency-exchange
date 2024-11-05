package app

import (
	"log"
	"net/http"

	"github.com/jvcosta-dev/go-currency-exchange/internal/app/routes"
)

func Server() {
	server := http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(),
	}

	log.Println("listening on :8080")
	server.ListenAndServe()
}
