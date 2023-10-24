.DEFAULT_GOAL := help

.PHONY: dev
dev: download db migrate-up ## start api server
	@air -c air.toml

.PHONY: db
db: ## start postgres
	@./result/bin/postgres.sh start

.PHONY: download
download: ## start postgres
	@go mod download

.PHONY: db-down
db-down: ## stop postgres
	@./result/bin/postgres.sh stop

.PHONY: seed
seed: db migrate-up ## stop postgres
	@go run cmd/seed/main.go

.PHONY: migrate-create
migrate-create: ## make migrate-create fileName=
	@migrate create -ext sql -dir db/migrations ${fileName}

.PHONY: migrate-up
migrate-up: ## run migration schema
	@migrate -path ./db/migrations -database "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down: ## rollback migration schema
	@migrate -path ./db/migrations -database "$(DATABASE_URL)" down

PHONY: sqlboiler-gen
sqlboiler-gen: ## generate code from schema
	@sqlboiler psql -c ./config/sqlboiler.toml


PHONY: api-schema-type-gen
api-schema-type-gen: ## generate response and request struct from openapi.yaml
	@oapi-codegen -generate types -package http openapi.yaml > internal/ports/http/types.gen.go

RED=\033[31m
GREEN=\033[32m
RESET=\033[0m

COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

PHONY: test
test: ## run all test
	@go test -v ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

.PHONY: help
help:
	@echo 'Usage: make [target]'
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
