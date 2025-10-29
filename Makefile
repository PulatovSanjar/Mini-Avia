COMPOSE ?= docker compose
ENV     ?= .env
MIGRATE ?= migrate

up: ## поднять контейнеры
	$(COMPOSE) --env-file $(ENV) up -d --build

down: ## остановить и удалить контейнеры
	$(COMPOSE) --env-file $(ENV) down --remove-orphans

init: ## полная инициализация: up + миграции
	$(MAKE) down
	$(MAKE) up