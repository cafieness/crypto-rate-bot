package scheduler_test

import (
	"context"
	"cryptobot/internal/mocks"
	"cryptobot/internal/scheduler"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdater_Update(t *testing.T) {

	mockRepo := &mocks.MockRateWriter{}

	updater := scheduler.NewUpdater(
		mockRepo,
		func(
			ctx context.Context,
			coin string,
		) (float64, error) {

			return 65000, nil
		},
		time.Minute,
	)

	updater.Update()

	assert.Equal(
		t,
		"ethereum",
		mockRepo.SavedCurrency,
	)

	assert.Equal(
		t,
		65000.0,
		mockRepo.SavedPrice,
	)
}
