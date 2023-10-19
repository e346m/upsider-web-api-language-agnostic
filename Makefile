.DEFAULT_GOAL := help

.PHONY: dev
dev: db ## start api server
	@echo "start"

.PHONY: db
db: ## start postgres
	@./result/bin/postgres.sh start

.PHONY: migrate-create
migrate-create: ## make migrate-create fileName=
	@migrate create -ext sql -dir db/migrations ${fileName}

.PHONY: migrate-up
migrate-up: ## run migration schema
	@migrate -path ./db/migrations -database $DATABASE_URL up

.PHONY: migrate-down
migrate-down: ## rollback migration schema
	@migrate -path ./db/migrations -database $DATABASE_URL down

.PHONY: help
help:
	@echo 'Usage: make [target]'
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
