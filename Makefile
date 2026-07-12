run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

test:
	go test ./... -race -count=1

test-cover:
	go test ./... -race -count=1 -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

docker-build:
	docker compose build

docker-up:
	docker compose up --build

docker-down:
	docker compose down

docker-psql:
	docker compose exec postgres psql -U postgres -d go_starter

.PHONY: run build test test-cover lint tidy docker-build docker-up docker-down docker-psql
