COMPOSE ?= docker compose
ENV     ?= .env
MIGRATE ?= migrate

up: ## поднять контейнеры
	$(COMPOSE) --env-file $(ENV) up -d --build

down: ## остановить и удалить контейнеры
	$(COMPOSE) --env-file $(ENV) down --remove-orphans

migrate-down: ## откатить одну миграцию (down 1)
	$(COMPOSE) --env-file $(ENV) run --rm $(MIGRATE) -path=/migrations -database $$DATABASE_URL down 1

init: ## полная инициализация: up + миграции
	$(MAKE) up