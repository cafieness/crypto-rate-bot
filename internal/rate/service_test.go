package rate_test

import (
	"context"
	"cryptobot/internal/rate"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRateRepository struct {
	mock.Mock
}

func (m *MockRateRepository) Save(
	ctx context.Context,
	currency string,
	price float64,
) error {

	args := m.Called(ctx, currency, price)

	return args.Error(0)
}

func (m *MockRateRepository) GetLatest(
	ctx context.Context,
	currency string,
) (rate.Rate, error) {

	args := m.Called(ctx, currency)

	return args.Get(0).(rate.Rate), args.Error(1)
}

func (m *MockRateRepository) GetDailyMinMax(
	ctx context.Context,
	currency string,
) (float64, float64, error) {

	args := m.Called(ctx, currency)

	return args.Get(0).(float64),
		args.Get(1).(float64),
		args.Error(2)
}

func (m *MockRateRepository) GetHourlyChange(
	ctx context.Context,
	currency string,
) (float64, error) {

	args := m.Called(ctx, currency)

	return args.Get(0).(float64), args.Error(1)
}

func TestRateService_GetLatest(t *testing.T) {
	repo := new(MockRateRepository)

	service := rate.NewRateService(repo)

	expected := rate.Rate{
		Currency: "bitcoin",
		Price:    65000,
	}

	repo.On(
		"GetLatest",
		mock.Anything,
		"bitcoin",
	).Return(
		expected,
		nil,
	)

	result, err := service.GetLatest(
		context.Background(),
		"bitcoin",
	)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	repo.AssertExpectations(t)
}

func TestRateService_Save(t *testing.T) {

	repo := new(MockRateRepository)

	service := rate.NewRateService(repo)

	repo.On(
		"Save",
		mock.Anything,
		"ethereum",
		2000.0,
	).Return(nil)

	err := service.Save(
		context.Background(),
		"ethereum",
		2000,
	)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestRateService_GetDailyMinMax(t *testing.T) {

	repo := new(MockRateRepository)

	s := rate.NewRateService(repo)

	repo.On(
		"GetDailyMinMax",
		mock.Anything,
		"bitcoin",
	).Return(
		60000.0,
		70000.0,
		nil,
	)

	min, max, err := s.GetDailyMinMax(
		context.Background(),
		"bitcoin",
	)

	assert.NoError(t, err)
	assert.Equal(t, 60000.0, min)
	assert.Equal(t, 70000.0, max)

	repo.AssertExpectations(t)
}

func TestRateService_GetHourlyChange(t *testing.T) {

	repo := new(MockRateRepository)

	s := rate.NewRateService(repo)

	repo.On(
		"GetHourlyChange",
		mock.Anything,
		"ethereum",
	).Return(
		2.5,
		nil,
	)

	change, err := s.GetHourlyChange(
		context.Background(),
		"ethereum",
	)

	assert.NoError(t, err)
	assert.Equal(t, 2.5, change)

	repo.AssertExpectations(t)
}
