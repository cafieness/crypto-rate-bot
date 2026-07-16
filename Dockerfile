FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o cryptobot ./cmd/app


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cryptobot .

COPY .env.example .env

EXPOSE 8081

CMD ["./cryptobot"]