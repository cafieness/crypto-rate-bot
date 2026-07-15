package service

import (
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

func (s *RateService) GetLatest(currency string) (domain.Rate, error) {
	return s.repo.GetLatest(currency)

}

func (s *RateService) GetDailyMinMax(currency string) (float64, float64, error) {
	return s.repo.GetDailyMinMax(currency)
}

func (s *RateService) GetHourlyChange(currency string) (float64, error) {
	return s.repo.GetHourlyChange(currency)
}

func (s *RateService) Save(currency string, price float64) error {
	return s.repo.Save(currency, price)
}
