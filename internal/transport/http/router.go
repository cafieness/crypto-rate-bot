package http

import (
	"cryptobot/internal/service"
	"net/http"
)

func NewRouter(rateService *service.RateService) http.Handler {
	mux := http.NewServeMux()
	handler := NewHandler(rateService)

	mux.HandleFunc(
		"GET /rates",
		handler.GetRates,
	)

	mux.HandleFunc(
		"GET /rates/",
		handler.GetRate,
	)

	return mux
}
