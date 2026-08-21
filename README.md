# Crypto Rate Bot

Service for tracking Bitcoin and Ethereum prices using CoinGeckoAPI, it stores rate history in Postgresql and sends scheduled updates through telegram bot

**Telegram bot:** [t.me/cafieness_crypto_rate_bot](https://t.me/cafieness_crypto_rate_bot)


## Features

- BTC and ETH rate tracking using CoingeckoAPI
- telegram bot with rate lookup and subscriptions
- postgresql storage for rates and subscriptions
- restapi for current crypto rates
- background updates
- swagger ui for documentation
- CI with tests, linting, and build checks
- autodeploy to oracle

## Architecture

```text
CoinGecko API
     |
     v
Rate Updater
     |
     v
PostgreSQL
     |
     +------> REST API
     |
     +------> Telegram Bot

Subscriptions
     |
     v
Notifier
     |
     v
Telegram
```

## Telegram Commands

```text
/start
/help
/rates
/rates bitcoin
/rates ethereum
/subscribe <currency> <minutes>
/unsubscribe
```

Example:

```text
/subscribe bitcoin 60
```

This sends Bitcoin rate updates every 60 minutes.

## REST API

Available endpoints:

```text
GET /rates
GET /rates/{currency}
GET /health
```

Examples:

```bash
curl http://localhost:8081/rates
curl http://localhost:8081/rates/bitcoin
curl http://localhost:8081/health
```

Swagger UI is available at:

```text
http://localhost:8082
```

## Run Locally

Clone the repository:

```bash
git clone https://github.com/cafieness/crypto-rate-bot.git
cd crypto-rate-bot
```

Create an environment file:

```bash
cp .env.example .env
```

Fill in the required environment variables:

```env
DATABASE_URL=
TELEGRAM_API_KEY=
COINGECKO_API_KEY=
COINGECKO_API_URL=
```

Start the application:

```bash
docker compose up -d
```

Check running services:

```bash
docker compose ps
```

View application logs:

```bash
docker compose logs -f app
```

Stop the application:

```bash
docker compose down
```

## Docker Image

The application Docker image is published to GitHub Container Registry:

```text
ghcr.io/cafieness/crypto-rate-bot:latest
```

Pull the latest image:

```bash
docker pull ghcr.io/cafieness/crypto-rate-bot:latest
```

## Development

Run tests:

```bash
go test ./...
```

Run tests with the race detector:

```bash
go test -race ./...
```

Run the linter:

```bash
golangci-lint run
```

Build the application:

```bash
go build ./...
```

## Notes

Service supports only rates for BTC and ETH as it was built mainly for my portfolio to practice building go architecture with external APIs, telegram bot and deploying to Oracle

Soo yeah feel free to try it out
