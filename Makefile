APP_NAME=pr-reviewer
PKG=github.com/quenyu/pr-reviewer

.PHONY: run build test lint docker-up docker-down

run:
	go run ./cmd/server

build:
	go build -o bin/$(APP_NAME) ./cmd/server

test:
	go test ./...

lint:
	golangci-lint run ./... || true

docker-up:
	docker compose up --build

docker-down:
	docker compose down

