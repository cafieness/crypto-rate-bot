package fetcher

import (
	"context"
	"cryptobot/internal/config"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	url        string
	key        string
	httpClient *http.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		url: cfg.CoinGeckoURL,
		key: cfg.CoinGeckoAPIKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) FetchRate(ctx context.Context, coin string) (float64, error) {
	params := url.Values{}

	params.Add("ids", coin)
	params.Add("vs_currencies", "usd")
	params.Add("x_cg_demo_api_key", c.key)

	requestURL := c.url + "?" + params.Encode()
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		requestURL,
		nil,
	)
	if err != nil {
		return 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Error("failed to close response body", "error", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("coingecko returned status %d: %s", resp.StatusCode, string(body))
	}

	slog.Info("coingecko response",
		"status", resp.StatusCode,
		"body", string(body),
		"url", requestURL,
	)

	var parsed CoinGeckoResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return 0, err
	}

	price, ok := parsed[coin]["usd"]
	if !ok {
		return 0, fmt.Errorf("price for %s not found in response: %s", coin, string(body))
	}

	return price, nil
}
