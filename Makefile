.PHONY: run-air run tidy fmt up down restart pull build clean logs logs-db logs-backend logs-frontend

# Local
run-air:
	cd backend && air

run:
	cd backend && go run main.go

tidy:
	cd backend && go mod tidy

fmt:
	cd backend && gofmt -w . && goimports -w .

# Docker Compose
up:
	docker compose up -d

up-db:
	docker compose up -d db

down:
	docker compose down

restart: down up

pull:
	docker compose pull

build:
	docker compose build

clean:
	docker compose down -t 0 -v

logs:
	docker compose logs -f

logs-db:
	docker compose logs -f db

logs-backend:
	docker compose logs -f backend

logs-frontend:
	docker compose logs -f frontend
