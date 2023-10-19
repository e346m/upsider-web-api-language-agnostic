.DEFAULT_GOAL := help

.PHONY: dev
dev: db ## start api server
	@echo "start"

.PHONEY: db
db: ## start postgres
	@./result/bin/postgres.sh start

.PHONY: migrate-create
migrate-create: ## make migrate-create fileName=
	@migrate create -ext sql -dir db/migrations ${fileName}

.PHONY: help
help:
	@echo 'Usage: make [target]'
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
