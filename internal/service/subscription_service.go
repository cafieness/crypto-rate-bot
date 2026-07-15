package service

import (
	"cryptobot/internal/domain"
	"fmt"
)

type SubscriptionService struct {
	repo domain.SubscriptionRepository
}

func NewSubscriptionService(
	repo domain.SubscriptionRepository,
) *SubscriptionService {
	return &SubscriptionService{
		repo: repo,
	}
}

func (s *SubscriptionService) Subscribe(chatID int64, minutes int, currency string) error {
	exists, err := s.repo.Exists(chatID, currency)

	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf(
			"Already subscribed to %s",
			currency,
		)
	}

	sub := domain.Subscription{
		ChatID:          chatID,
		IntervalMinutes: minutes,
		Currency:        currency,
		Active:          true,
	}
	return s.repo.Create(sub)
}

func (s *SubscriptionService) Unsubscribe(chatID int64) error {
	return s.repo.Deactivate(chatID)
}

func (s *SubscriptionService) GetActive() ([]domain.Subscription, error) {
	return s.repo.GetActive()
}

func (s *SubscriptionService) UpdateLastSent(
	chatID int64,
	currency string,
) error {
	return s.repo.UpdateLastSent(chatID, currency)
}
