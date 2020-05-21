DOCKER_COMPOSE = docker-compose
TESTS_DIR = tests

up:
	$(DOCKER_COMPOSE) up -d --remove-orphans

logs:
	$(DOCKER_COMPOSE) logs -f


down:
	$(DOCKER_COMPOSE) down

test:
	$(DOCKER_COMPOSE) run --rm notesapp.e2e go test ./$(TESTS_DIR)/...