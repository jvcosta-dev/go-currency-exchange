package routes

import (
	"net/http"

	"github.com/jvcosta-dev/go-currency-exchange/internal/app/handlers"
)

func Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /convert/{from}/{to}/{amount}", handlers.Convert)
	mux.HandleFunc("GET /rates", handlers.Rates)
	mux.HandleFunc("GET /latest/{base}", handlers.Rate)

	return mux
}
