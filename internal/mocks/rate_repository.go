package mocks

import (
	"context"
	"cryptobot/internal/domain"
)

type MockRateRepository struct {
	SaveFunc func(
		ctx context.Context,
		currency string,
		price float64,
	) error

	GetLatestFunc func(
		ctx context.Context,
		currency string,
	) (domain.Rate, error)

	GetDailyMinMaxFunc func(
		ctx context.Context,
		currency string,
	) (float64, float64, error)

	GetHourlyChangeFunc func(
		ctx context.Context,
		currency string,
	) (float64, error)
}

func (m *MockRateRepository) Save(
	ctx context.Context,
	currency string,
	price float64,
) error {
	return m.SaveFunc(ctx, currency, price)
}

func (m *MockRateRepository) GetLatest(
	ctx context.Context,
	currency string,
) (domain.Rate, error) {
	return m.GetLatestFunc(ctx, currency)
}

func (m *MockRateRepository) GetDailyMinMax(
	ctx context.Context,
	currency string,
) (float64, float64, error) {
	return m.GetDailyMinMaxFunc(ctx, currency)
}

func (m *MockRateRepository) GetHourlyChange(
	ctx context.Context,
	currency string,
) (float64, error) {
	return m.GetHourlyChangeFunc(ctx, currency)
}
