##---------- Preliminaries ----------------------------------------------------
.POSIX:     # Get reliable POSIX behaviour
.SUFFIXES:  # Clear built-in inference rules

##---------- Variables --------------------------------------------------------

##---------- Export .env as vars ----------------------------------------------
include .env
export

help: ## Show this help message (default)
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

##---------- APPLICATION -------------------------------------------------------

run: ## Run application
	air

lint: ## Lint code
	golangci-lint run --enable-all

test: ## Test the application
	go clean -testcache
	gotest -v ./...

##---------- UTILITY ----------------------------------------------------------

compose: ## Run docker compose stack
	docker-compose rm -f
	docker-compose up

deploy: ## Deploy
	KO_DOCKER_REPO=$$KO_DOCKER_REPO ko resolve -f kubernetes-application.yaml | kubectl apply -f -

sqlgen: ## Generate sql via sqlc
	sqlc generate

cssgen: ## Generate css
	bun run css

update: ## Update all dependencies
	go get -u
	go mod tidy

.PHONY: test-coverage test
