DOCKER_COMPOSE = docker-compose
TESTS_DIR = tests
DB_CONNECTION = postgres://db-user:db-pass@db:5432/db-name?sslmode=disable
TEST_DB_CONNECTION = postgres://db-user:db-pass@db-test:5432/db-name?sslmode=disable
RUN_MIGRATION = run --rm migrate -source file://migrations -database

up:
	$(DOCKER_COMPOSE) up -d --remove-orphans

logs:
	$(DOCKER_COMPOSE) logs -f

down:
	$(DOCKER_COMPOSE) down

test:
	$(DOCKER_COMPOSE) run --rm notesapp.e2e go test ./$(TESTS_DIR)/...

migrate-db-up:
	$(DOCKER_COMPOSE) $(RUN_MIGRATION) $(DB_CONNECTION) up
	$(DOCKER_COMPOSE) $(RUN_MIGRATION) $(TEST_DB_CONNECTION) up

migrate-db-down:
	$(DOCKER_COMPOSE) $(RUN_MIGRATION) $(DB_CONNECTION) down 1
	$(DOCKER_COMPOSE) $(RUN_MIGRATION) $(TEST_DB_CONNECTION) down 1