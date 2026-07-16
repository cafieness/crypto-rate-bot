package service

import (
	"context"
	"cryptobot/internal/domain"
)

type RateService struct {
	repo domain.RateRepository
}

func NewRateService(
	repo domain.RateRepository,
) *RateService {
	return &RateService{
		repo: repo,
	}
}

func (s *RateService) GetLatest(ctx context.Context, currency string) (domain.Rate, error) {
	return s.repo.GetLatest(ctx, currency)

}

func (s *RateService) GetDailyMinMax(ctx context.Context, currency string) (float64, float64, error) {
	return s.repo.GetDailyMinMax(ctx, currency)
}

func (s *RateService) GetHourlyChange(ctx context.Context, currency string) (float64, error) {
	return s.repo.GetHourlyChange(ctx, currency)
}

func (s *RateService) Save(ctx context.Context, currency string, price float64) error {
	return s.repo.Save(ctx, currency, price)
}
