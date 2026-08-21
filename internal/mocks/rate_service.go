package mocks

import (
	"context"
	"cryptobot/internal/rate"
)

type MockService struct {
	RateResult rate.Rate
	Err        error
}

func (m *MockService) GetLatest(
	ctx context.Context,
	currency string,
) (rate.Rate, error) {

	return m.RateResult, m.Err
}
