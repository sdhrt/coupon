FROM golang:1.23.5-bullseye AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./build/main ./cmd/api

FROM debian:bookworm

RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    && rm -rf /var/lib/apt/lists/*

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz \
  | tar xvz -C /usr/local/bin \
  && chmod +x /usr/local/bin/migrate

WORKDIR /app

COPY --from=builder /app/build/main /app/main

COPY --from=builder /app/migration /app/migration

ENTRYPOINT ["/app/main"]
