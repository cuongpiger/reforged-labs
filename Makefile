GO_CMD ?= "/Volumes/veronica/git-cuongpiger/go-env/go-1.24.1/go1.24.1/bin/go"
CURDIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GIN_MODE ?= release
APP_ENV ?= development

API_SERVICE_CONFIG_FILE ?= $(CURDIR)/env/api-service-config-file.yaml


.PHONY: run-api-service
run-api-service:
	@echo "Running API service at localhost"
	export APP_ENV=$(APP_ENV) && \
		export GIN_MODE=$(GIN_MODE) && \
		$(GO_CMD) run cmd/api-service/main.go --config-file=$(API_SERVICE_CONFIG_FILE)
