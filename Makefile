SHELL := /bin/bash
export PATH := $(shell go env GOPATH)/bin:${PATH}
.DEFAULT_GOAL := all
.PHONY: all
all: ## build pipeline
all: mod gen build spell lint test

.PHONY: precommit
precommit: ## validate the branch before commit
precommit: all vuln

.PHONY: ci
ci: ## CI build pipeline
ci: precommit diff

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## remove files created during build pipeline
	$(call print-target)
	rm -rf dist
	rm -f coverage.*
	rm -f '"$(shell go env GOCACHE)/../golangci-lint"'
	go clean -i -cache -testcache -modcache -fuzzcache -x

.PHONY: mod
mod: ## go mod tidy
	$(call print-target)
	echo ${PATH}
	go mod tidy

.PHONY: gen
gen: ## go generate
	$(call print-target)
	go generate ./...

.PHONY: build
build: ## goreleaser build
	$(call print-target)
	go tool goreleaser build --clean --single-target --snapshot

.PHONY: spell
spell: ## misspell
	$(call print-target)
	go tool misspell -error -locale=US -w **.md

.PHONY: lint
lint: ## golangci-lint
	$(call print-target)
	go tool golangci-lint run --fix

.PHONY: vuln
vuln: ## govulncheck
	$(call print-target)
	go tool govulncheck ./...

.PHONY: test
test: ## go test
	$(call print-target)
	go test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: diff
diff: ## git diff
	$(call print-target)
	git diff --exit-code
	RES=$$(git status --porcelain) ; if [ -n "$$RES" ]; then echo $$RES && exit 1 ; fi

config/local/grafana:
	mkdir -p config/local/grafana

.PHONY: start-otel-collector
start-otel-collector: config/local/grafana config/local/prometheus.yml config/local/otel-collector-config.yaml
	USERID=$(shell id -u) GROUPID=$(shell id -g) docker compose up

.PHONY: stop-otel-collector
stop-otel-collector:
	docker compose down

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
