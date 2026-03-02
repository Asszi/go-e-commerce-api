.PHONY: help build run dev lint migrate-up migrate-down docker-up docker-down

help:
	@echo "Available commands:"
	@echo " make build          - Builds the application"
	@echo " make run            - Run the application"
	@echo " make dev            - Run the application in development mode"
	@echo " make lint           - Run linter on the codebase"
	@echo " make migrate-up     - Apply database migrations"
	@echo " make migrate-down   - Rollback database migrations"
	@echo " make docker-up      - Run docker compose up"
	@echo " make docker-down    - Run docker compose down"

build:
	go build -o bin/app ./cmd/api

run:
	go run ./cmd/api

dev:
	go run ./cmd/api

lint:
	golangci-lint run ./...

migrate-up:
	migrate -source file://db/migrations -database "postgresql://postgres:password@localhost:5432/ecommerce_shop?sslmode=disable" up

migrate-down:
	migrate -source file://db/migrations -database "postgresql://postgres:password@localhost:5432/ecommerce_shop?sslmode=disable" down

docker-up:
	podman compose -f docker/docker-compose.yml up -d

docker-down:
	podman compose -f docker/docker-compose.yml down
