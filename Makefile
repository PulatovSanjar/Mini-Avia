COMPOSE ?= docker compose
ENV     ?= .env
MIGRATE ?= migrate

up: ## поднять контейнеры
	$(COMPOSE) --env-file $(ENV) up -d --build

down: ## остановить и удалить контейнеры
	$(COMPOSE) --env-file $(ENV) down --remove-orphans

migrate-down:
	migrate -path internal/migrations -database "postgres://miniavia:miniavia@localhost:5432/miniavia?sslmode=disable" down -all

init: ## полная инициализация: up + миграции
	$(MAKE) down
	$(MAKE) up