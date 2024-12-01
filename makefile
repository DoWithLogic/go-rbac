IN_PROGRESS = "is in progress ..."

# Define the migration directory
MIGRATION_DIR = migrations/files

# The database connection string (can be set via environment variables or override here)
DB_CONN = "root:pwd@tcp(localhost:3306)/go_rbac"

## help: prints this help message
.PHONY: help
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## mod: will pull all dependency
.PHONY: mod
mod:
	@echo "make mod ${IS_IN_PROGRESS}"
	@rm -rf ./vendor ./go.sum
	@go mod tidy
	@go mod vendor

## install: will install migration and golangci-lint
.PHONY: install
install:
	@echo "make install ${IS_IN_PROGRESS}"
	@go install go.uber.org/mock/mockgen@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/pressly/goose/v3/cmd/goose@latest


## setup: Set up infrastructure dependencies for local environtment
.PHONY: setup
setup:
	@echo "make setup ${IS_IN_PROGRESS}"
	@docker-compose -f ./infrastructure/docker-compose.yml up -d
	@sleep 8

## down: Set down infrastructure dependencies for local environtment
.PHONY: down
down: 
	@echo "make down ${IS_IN_PROGRESS}"
	@docker-compose -f ./infrastructure/docker-compose.yml down -t 1

## migrate-create-file: run create file migration scripts.
.PHONY: migrate-create-file
migrate-create-file:
	@goose -dir migrations/files/ create $(FILENAME) sql

# Migrate up one migration at a time
.PHONY: migrate-up
migrate-up:
	@echo "Running migrate up"
	@goose -dir $(MIGRATION_DIR) mysql $(DB_CONN) up

# Migrate down one migration at a time
.PHONY: migrate-down
migrate-down:
	@echo "Running migrate down"
	@goose -dir $(MIGRATION_DIR) mysql $(DB_CONN) down 1

# Migrate reset (roll back all migrations)
.PHONY: migrate-reset
migrate-reset:
	@echo "Rolling back all migrations"
	@goose -dir $(MIGRATION_DIR) mysql $(DB_CONN) reset

# Migrate status (optional - to check migration status)
.PHONY: migrate-status
migrate-status:
	@echo "Checking migration status"
	@goose -dir $(MIGRATION_DIR) mysql $(DB_CONN) status

## mock-repo: will generate mock for repositories interfaces
.PHONY: mock-repo
mock-repo:
	@echo "make mock-repo ${IS_IN_PROGRESS}"
	@mockgen -source internal/${DOMAIN}/repository.go -destination mocks/${DOMAIN}/repository_mock.go -package=mocks

## mock-uc: will generate mock for usecase interfaces
.PHONY: mock-uc
mock-uc:
	@echo "make mock-uc ${IS_IN_PROGRESS}"
	@mockgen -source internal/${DOMAIN}/usecase.go -destination mocks/${DOMAIN}/usecase_mock.go -package=mocks

## mock-pkg: will generate mock for pkg interfaces
.PHONY: mock-pkg
mock-pkg:
	@echo "make mock-pkg ${IS_IN_PROGRESS}"
	@mockgen -source pkg/${DOMAIN}/${DOMAIN}.go -destination mocks/${DOMAIN}/${DOMAIN}_mock.go -package=mocks


# Run: Run the application with migrations applied
.PHONY: run
run: setup migrate-up mod install
	@echo "Starting the application"
	@go run cmd/main.go

## swagger-gen: generate swagger documentation
.PHONY: swagger-gen
swagger-gen:
	@swag init -g cmd/main.go -md ./docs/ && ./docs/fix.sh