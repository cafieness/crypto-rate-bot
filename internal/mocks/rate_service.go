package mocks

import (
	"context"
	"cryptobot/internal/domain"
)

type MockRateService struct {
	RateResult domain.Rate
	Err        error
}

func (m *MockRateService) GetLatest(
	ctx context.Context,
	currency string,
) (domain.Rate, error) {

	return m.RateResult, m.Err
}
