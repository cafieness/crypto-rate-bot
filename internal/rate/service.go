package rate

import (
	"context"
)

type RateService struct {
	repo RateRepository
}

func NewRateService(
	repo RateRepository,
) *RateService {
	return &RateService{
		repo: repo,
	}
}

func (s *RateService) GetLatest(ctx context.Context, currency string) (Rate, error) {
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
