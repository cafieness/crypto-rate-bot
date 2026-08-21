package http

import (
	"cryptobot/internal/rate"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

type Handler struct {
	rateService rate.Service
}

func NewHandler(rateService rate.Service) *Handler {
	return &Handler{
		rateService: rateService,
	}
}

func (h *Handler) GetRate(
	w http.ResponseWriter,
	r *http.Request,
) {
	currency := strings.TrimPrefix(
		r.URL.Path,
		"/rates/",
	)
	if currency == "" {
		http.Error(
			w,
			"currency is required",
			http.StatusBadRequest,
		)
		return
	}

	rate, err := h.rateService.GetLatest(
		r.Context(),
		currency,
	)

	if err != nil {
		http.Error(
			w,
			"rate not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	err = json.NewEncoder(w).Encode(rate)

	if err != nil {
		http.Error(
			w,
			"failed to encode response",
			http.StatusInternalServerError,
		)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	}); err != nil {
		slog.Error("failed to encode health response", "error", err)
	}
}

func (h *Handler) GetRates(
	w http.ResponseWriter,
	r *http.Request,
) {

	currencies := []string{
		"bitcoin",
		"ethereum",
	}

	var rates []rate.Rate

	for _, currency := range currencies {

		rate, err := h.rateService.GetLatest(r.Context(), currency)

		if err != nil {
			http.Error(
				w,
				"failed to get rates",
				http.StatusInternalServerError,
			)
			return
		}

		rates = append(
			rates,
			rate,
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	if err := json.NewEncoder(w).Encode(rates); err != nil {
		http.Error(
			w,
			"failed to encode response",
			http.StatusInternalServerError,
		)
	}
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS",
		)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewRouter(rateService *rate.RateService) http.Handler {
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
	mux.HandleFunc(
		"GET /health",
		HealthHandler,
	)

	return mux
}
