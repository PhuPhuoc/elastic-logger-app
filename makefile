.DEFAULT_GOAL := help

# App info
APP_NAME := elastic-logger-app
BIN_DIR := bin
MAIN_FILE := cmd/main.go
SERVICE_NAME := elastic-logger-app
DOCKER_OWNER := phuoctran
IMAGE_VER := v1 

# Debug flags
GCFLAGS := all=-N -l

.PHONY: help build run build-debug debug up down docker-build docker-tag docker-push clean

help: ## Show all available commands
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*##"}; {printf "  \033[32m%-15s\033[0m %s\n", $$1, $$2}'

# =====================================================
# Go build
# =====================================================

build: ## Build the Go application
	@echo "==> Building $(APP_NAME)"
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_FILE)

run: build ## Run the application
	./$(BIN_DIR)/$(APP_NAME)

build-debug: ## Build the app in debug mode
	@echo "==> Building debug binary"
	mkdir -p $(BIN_DIR)
	go build -gcflags="$(GCFLAGS)" -o $(BIN_DIR)/$(APP_NAME)-debug $(MAIN_FILE)

debug: build-debug ## Run the app with delve debugger
	@command -v dlv >/dev/null 2>&1 || go install github.com/go-delve/delve/cmd/dlv@latest
	dlv exec ./$(BIN_DIR)/$(APP_NAME)-debug

# =====================================================
# Docker
# =====================================================

docker-build: ## Build docker image
	docker build -t $(SERVICE_NAME):$(IMAGE_VER) .

docker-tag: ## Tag docker image
	docker tag $(SERVICE_NAME):$(IMAGE_VER) $(DOCKER_OWNER)/$(SERVICE_NAME):$(IMAGE_VER)

docker-push: ## Push docker image
	docker push $(DOCKER_OWNER)/$(SERVICE_NAME):$(IMAGE_VER)

# =====================================================
# Docker Compose
# =====================================================

up: ## Start docker-compose
	docker compose up -d

down: ## Stop docker-compose
	docker compose down

# =====================================================
# Utility
# =====================================================

clean: ## Clean build artifacts
	rm -rf $(BIN_DIR)


# =====================================================
# Goose Migration
# =====================================================

GOOSE_DRIVER := mysql
GOOSE_DBSTRING := "root:rootpwd@tcp(localhost:3306)/cmd_elastic_app?parseTime=true&multiStatements=true"
MIGRATION_DIR := migration

gs-status: ## Show migration status
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(MIGRATION_DIR) status

gs-up: ## Run all migrations
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(MIGRATION_DIR) up

gs-down: ## Roll back last migration
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(MIGRATION_DIR) down

gs-create: ## Create a new migration file: make migrate-create name=add_table
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(MIGRATION_DIR) create $(name) sql

gs-create-seed: ## Create a new seed file: make migrate-create-seed name=seed_users
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(MIGRATION_DIR) create $(name)_seed sql

gs-seed: ## Apply only seed migrations (file name contains "_seed")
	@echo "==> Running only seed migrations…"
	@for file in $(MIGRATION_DIR)/*_seed.sql; do \
		echo "Applying $$file"; \
		GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(MIGRATION_DIR) up-to $$(basename $$file | cut -d'_' -f1); \
	done



# =====================================================
# sqlc generate 
# =====================================================
gen-account: ## generate sqlc for module account 
	cd modules/account/infras/commandrepo && sqlc generate
