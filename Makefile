#!make
include .env
export $(shell sed 's/=.*//' .env)

build:
	go mod download\
	&& GOOS=linux CGO_ENABLED=0 go build -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up --remove-orphans app

swag:
	swag init -g internal/app/app.go

test:
	go test ./...

lint:
	golangci-lint run

postgres_url ?= "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

migrate-up:
	migrate -path database/migrations/postgres -database ${postgres_url} -verbose up

migrate-down:
	migrate -path database/migrations/postgres -database ${postgres_url} -verbose down