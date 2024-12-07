# Этап сборки
FROM golang:1.23.3 AS builder

WORKDIR /rest-app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# Этап запуска
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /rest-app

COPY --from=builder /rest-app/main .

EXPOSE 8080

CMD ["./main"]
