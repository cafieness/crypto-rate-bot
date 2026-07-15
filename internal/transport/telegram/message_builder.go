package telegram

import (
	"cryptobot/internal/service"
	"fmt"
	"strings"
)

func BuildRateMessage(currency string, rateService *service.RateService) (string, error) {
	rate, err := rateService.GetLatest(currency)

	if err != nil {
		return "", err
	}

	min, max, err := rateService.GetDailyMinMax(currency)

	if err != nil {
		return "", err
	}

	change, err := rateService.GetHourlyChange(currency)

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
