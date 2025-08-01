#!make
include .env.local

build:
	@go build -o ./build/main ./cmd/api

run: build
	@./build/main

help: build 
	./build/main -h

dryrun:
	go run ./cmd/api

migrate-up: 
	migrate -database ${POSTGRES_DSN} -path ./migration up

migrate-down: 
	migrate -database ${POSTGRES_DSN} -path ./migration down
