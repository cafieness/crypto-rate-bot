APP_NAME=cryptobot

.PHONY: run test coverage lint build docker-up docker-down clean


run:
	go run ./cmd/app


test:
	go test ./...


coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out


lint:
	golangci-lint run


build:
	go build -o bin/$(APP_NAME) ./cmd/app


docker-up:
	docker compose up --build


docker-down:
	docker compose down


clean:
	rm -rf bin coverage.out