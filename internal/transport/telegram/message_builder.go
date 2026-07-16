package telegram

import (
	"context"
	"cryptobot/internal/service"
	"fmt"
	"strings"
)

func BuildRateMessage(ctx context.Context, currency string, rateService *service.RateService) (string, error) {
	rate, err := rateService.GetLatest(ctx, currency)

	if err != nil {
		return "", err
	}

	min, max, err := rateService.GetDailyMinMax(ctx, currency)

	if err != nil {
		return "", err
	}

	change, err := rateService.GetHourlyChange(ctx, currency)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		`%s
		Current rate: %.2f$
		Daily min: %.2f$
		Daily max: %.2f$
		Hourly change: %.2f%%`,
		strings.ToUpper(currency),
		rate.Price,
		min,
		max,
		change,
	), nil
}
