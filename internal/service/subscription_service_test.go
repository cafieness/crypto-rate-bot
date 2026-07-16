package service_test

import (
	"context"
	"cryptobot/internal/domain"
	"cryptobot/internal/mocks"
	"cryptobot/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscriptionService_Subscribe(t *testing.T) {
	repo := &mocks.MockSubscriptionRepository{}

	service := service.NewSubscriptionService(repo)

	err := service.Subscribe(
		context.Background(),
		12345,
		60,
		"bitcoin",
	)

	assert.NoError(t, err)

	assert.Len(
		t,
		repo.Subscriptions,
		1,
	)

	assert.Equal(
		t,
		"bitcoin",
		repo.Subscriptions[0].Currency,
	)
}

func TestSubscriptionService_SubscribeAlreadyExists(t *testing.T) {

	repo := &mocks.MockSubscriptionRepository{
		ExistsResult: true,
	}

	service := service.NewSubscriptionService(repo)

	err := service.Subscribe(
		context.Background(),
		123,
		30,
		"ethereum",
	)

	assert.Error(t, err)
}

func TestSubscriptionService_Unsubscribe(t *testing.T) {

	repo := &mocks.MockSubscriptionRepository{
		Subscriptions: []domain.Subscription{
			{
				ChatID:   123,
				Currency: "bitcoin",
				Active:   true,
			},
		},
	}

	service := service.NewSubscriptionService(repo)

	err := service.Unsubscribe(
		context.Background(),
		123,
	)

	assert.NoError(t, err)

	assert.False(
		t,
		repo.Subscriptions[0].Active,
	)

}

func TestSubscriptionService_GetActive(t *testing.T) {

	repo := &mocks.MockSubscriptionRepository{
		Subscriptions: []domain.Subscription{
			{
				ChatID:   1,
				Currency: "bitcoin",
				Active:   true,
			},
		},
	}

	service := service.NewSubscriptionService(repo)

	result, err := service.GetActive(
		context.Background(),
	)

	assert.NoError(t, err)

	assert.Len(
		t,
		result,
		1,
	)
}
