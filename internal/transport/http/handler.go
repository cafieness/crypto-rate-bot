package http

import (
	"cryptobot/internal/domain"
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	rateService domain.RateService
}

func NewHandler(rateService domain.RateService) *Handler {
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

func (h *Handler) GetRates(
	w http.ResponseWriter,
	r *http.Request,
) {

	currencies := []string{
		"bitcoin",
		"ethereum",
	}

	var rates []domain.Rate

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
