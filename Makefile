# Variables
COMPOSE_FILE = docker/docker-compose.yml

# Targets
up:
	docker-compose -f $(COMPOSE_FILE) up -d

down:
	docker-compose -f $(COMPOSE_FILE) down

restart:
	docker-compose -f $(COMPOSE_FILE) down && docker-compose -f $(COMPOSE_FILE) up -d

logs:
	docker-compose -f $(COMPOSE_FILE) logs -f

ps:
	docker-compose -f $(COMPOSE_FILE) ps

clean:
	docker-compose -f $(COMPOSE_FILE) down --volumes --remove-orphans
