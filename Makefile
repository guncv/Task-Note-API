DC = docker compose
FILE = -f docker-compose.local.yml
NETWORK_NAME = task-network

.PHONY: run stop restart build ps clean rebuild exec ensure-network test mock-gen

ensure-network:
	@echo "🔍 Checking if Docker network '$(NETWORK_NAME)' exists..."
	@if ! docker network ls --format '{{.Name}}' | grep -wq "$(NETWORK_NAME)"; then \
		echo "⛓️  Creating network '$(NETWORK_NAME)'..."; \
		docker network create $(NETWORK_NAME); \
	else \
		echo "✅ Network '$(NETWORK_NAME)' already exists."; \
	fi

run: ensure-network
	$(DC) $(FILE) up

stop:
	$(DC) $(FILE) down

build: ensure-network
	$(DC) $(FILE) build

clean:
	$(DC) $(FILE) down --volumes --remove-orphans

reset: clean ensure-network
	$(DC) $(FILE) build --no-cache
	$(DC) $(FILE) up

mock-gen:
	mockery --all

swagger-gen:
	swag init

test:
	go test -v -cover ./services & go test -v -cover ./utils