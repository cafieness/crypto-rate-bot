package coingecko_test

import (
	"context"
	"cryptobot/internal/coingecko"
	"cryptobot/internal/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchRate_Success(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			_, err := w.Write([]byte(`
			{
				"bitcoin": {
					"usd": 65000
				}
			}
			`))

			assert.NoError(t, err)
		}),
	)

	defer server.Close()

	cfg := &config.Config{
		CoinGeckoURL:    server.URL,
		CoinGeckoAPIKey: "test",
	}

	client := coingecko.NewClient(cfg)

	price, err := client.FetchRate(
		context.Background(),
		"bitcoin",
	)

	assert.NoError(t, err)

	assert.Equal(
		t,
		65000.0,
		price,
	)
}
