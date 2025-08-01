#!make
include .env.local

build:
	go build -o ./dist/main ./cmd/api

run: build
	./dist/main

help: build 
	./dist/main -h

dryrun:
	go run ./cmd/api

migrate-up: 
	migrate -database ${POSTGRES_DSN} -path ./migration up

migrate-down: 
	migrate -database ${POSTGRES_DSN} -path ./migration down
