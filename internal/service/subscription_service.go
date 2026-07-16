package service

import (
	"context"
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

func (s *SubscriptionService) Subscribe(ctx context.Context, chatID int64, minutes int, currency string) error {
	exists, err := s.repo.Exists(ctx, chatID, currency)

	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf(
			"already subscribed to %s",
			currency,
		)
	}

	sub := domain.Subscription{
		ChatID:          chatID,
		IntervalMinutes: minutes,
		Currency:        currency,
		Active:          true,
	}
	return s.repo.Create(ctx, sub)
}

func (s *SubscriptionService) Unsubscribe(ctx context.Context, chatID int64) error {
	return s.repo.Deactivate(ctx, chatID)
}

func (s *SubscriptionService) GetActive(ctx context.Context) ([]domain.Subscription, error) {
	return s.repo.GetActive(ctx)
}

func (s *SubscriptionService) UpdateLastSent(
	ctx context.Context,
	chatID int64,
	currency string,
) error {
	return s.repo.UpdateLastSent(ctx, chatID, currency)
}
