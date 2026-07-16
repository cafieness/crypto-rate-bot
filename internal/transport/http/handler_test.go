package http_test

import (
	"cryptobot/internal/domain"
	"cryptobot/internal/mocks"
	httpTransport "cryptobot/internal/transport/http"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_GetRate(t *testing.T) {

	mockService := &mocks.MockRateService{
		RateResult: domain.Rate{
			Currency: "bitcoin",
			Price:    60000,
		},
	}

	handler := httpTransport.NewHandler(mockService)

	req := httptest.NewRequest(
		http.MethodGet,
		"/rates/bitcoin",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.GetRate(
		recorder,
		req,
	)

	assert.Equal(
		t,
		http.StatusOK,
		recorder.Code,
	)

	assert.Contains(
		t,
		recorder.Body.String(),
		"bitcoin",
	)

	assert.Contains(
		t,
		recorder.Body.String(),
		"60000",
	)
}

func TestHandler_GetRate_NotFound(t *testing.T) {

	mockService := &mocks.MockRateService{
		Err: errors.New("rate not found"),
	}

	handler := httpTransport.NewHandler(mockService)

	req := httptest.NewRequest(
		http.MethodGet,
		"/rates/bitcoin",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.GetRate(
		recorder,
		req,
	)

	assert.Equal(
		t,
		http.StatusNotFound,
		recorder.Code,
	)
}
