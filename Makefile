.DEFAULT_GOAL := help

.PHONY: build
build: proto
	go build -o bin/gowallet

.PHONY: test
test: build
	go test -v ./...

.PHONY: lint
lint: build
	golangci-lint run ./...

.PHONY: build-docker
build-docker:
	docker build . -t gowallet

.PHONY: run-docker
build-docker:
	docker-compose up -d

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Available targets:"
	@echo "  build			build the application"
	@echo "  build-docker	build docker image"
	@echo "  run-docker		build and run in docker"
	@echo "  test			run the tests"
	@echo "  help			display this help"
	@echo ""
