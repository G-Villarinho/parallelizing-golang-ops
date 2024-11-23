DOCKER_COMPOSE_FILE = docker-compose.yml
APP_NAME = parallelizing-app
MAIN_FILE = main.go

.PHONY: docker-up docker-down run-app docker-clean start docker-rebuild

docker-up:
	@echo "Subindo os serviços do Docker..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "Serviços do Docker estão rodando."

docker-down:
	@echo "Derrubando os serviços do Docker..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down
	@echo "Serviços do Docker foram parados."

run-app:
	@echo "Iniciando a aplicação Go..."
	go run $(MAIN_FILE)

docker-clean:
	@echo "Removendo os contêineres e volumes..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --volumes --remove-orphans

start:
	@echo "Iniciando os serviços do Docker e a aplicação Go..."
	make docker-up
	sleep 5 # Aguarda os serviços do Docker inicializarem
	make run-app


docker-rebuild:
	@echo "Rebuildando os serviços do Docker..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build -d
