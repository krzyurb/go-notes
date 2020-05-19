DOCKER_COMPOSE = docker-compose
GO_COMMAND = go
TESTS_DIR = tests

up:
	$(DOCKER_COMPOSE) up -d --remove-orphans

logs:
	$(DOCKER_COMPOSE) logs -f


down:
	$(DOCKER_COMPOSE) down

test:
	$(GO_COMMAND) test ./$(TESTS_DIR)/...