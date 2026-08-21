package mocks

import "context"

type MockRateWriter struct {
	SavedCurrency string
	SavedPrice    float64
	Err           error
}

func (m *MockRateWriter) Save(
	ctx context.Context,
	currency string,
	price float64,
) error {

	m.SavedCurrency = currency
	m.SavedPrice = price

	return m.Err
}
