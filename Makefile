GO_CMD ?= "/Volumes/veronica/git-cuongpiger/go-env/go-1.24.1/go1.24.1/bin/go"

CURDIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GIN_MODE ?= release
APP_ENV ?= development

PRE_RUN_API_SERVICE_FILE ?= $(CURDIR)/hack/pre-run-api-service
API_SERVICE_CONFIG_FILE ?= $(CURDIR)/hack/api-service-config-file.yaml


.PHONY: run-api-service
run-api-service:
	@echo "Running API service at localhost"
	source $(PRE_RUN_API_SERVICE_FILE) && $(GO_CMD) run cmd/api-service/main.go --config-file=$(API_SERVICE_CONFIG_FILE)
