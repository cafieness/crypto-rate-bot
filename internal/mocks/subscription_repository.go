package mocks

import (
	"context"
	"cryptobot/internal/subscription"
)

type MockSubscriptionRepository struct {
	Subscriptions []subscription.Subscription

	CreateErr    error
	ExistsResult bool
	ExistsErr    error
}

func (m *MockSubscriptionRepository) Create(
	ctx context.Context,
	sub subscription.Subscription,
) error {

	if m.CreateErr != nil {
		return m.CreateErr
	}

	m.Subscriptions = append(
		m.Subscriptions,
		sub,
	)

	return nil
}

func (m *MockSubscriptionRepository) Exists(
	ctx context.Context,
	chatID int64,
	currency string,
) (bool, error) {

	return m.ExistsResult, m.ExistsErr
}

func (m *MockSubscriptionRepository) Deactivate(
	ctx context.Context,
	chatID int64,
) error {

	for i := range m.Subscriptions {
		if m.Subscriptions[i].ChatID == chatID {
			m.Subscriptions[i].Active = false
		}
	}

	return nil
}

func (m *MockSubscriptionRepository) GetActive(
	ctx context.Context,
) ([]subscription.Subscription, error) {

	return m.Subscriptions, nil
}

func (m *MockSubscriptionRepository) UpdateLastSent(
	ctx context.Context,
	chatID int64,
	currency string,
) error {

	return nil
}
