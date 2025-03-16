
# CMD bin path for MacOS machine
GO_CMD ?= "/Volumes/veronica/git-cuongpiger/go-env/go-1.24.1/go1.24.1/bin/go"
DOCKER_CMD ?= "/usr/local/bin/docker"
TAG ?= v0.0.0
REGISTRY ?= "docker.io/manhcuong8499"
CONTROLLER_IMG_TAG ?= "$(REGISTRY)/reforged-labs-api-service:$(TAG)"


# CMD bin path for Linux machine
# GO_CMD ?= "/usr/local/go/bin/go"
# DOCKER_CMD ?= "/usr/bin/docker"

CURDIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
ENV_FILE ?= $(CURDIR)/hack/env
API_SERVICE_CONFIG_FILE ?= $(CURDIR)/hack/api-service-config-file.yaml
# Set build time variables including version details
LDFLAGS := $(shell source ./hack/version.sh; version::ldflags)
ARCH ?= amd64

include $(ENV_FILE)


.PHONY: check-ldflags
check-ldflags:
	@echo $(LDFLAGS)


.PHONY: run-api-service
run-api-service:
	@echo "Running API service at localhost"
	export APP_ENV=$(APP_ENV) && \
		export GIN_MODE=$(GIN_MODE) && \
		$(GO_CMD) run cmd/api-service/main.go --config-file=$(API_SERVICE_CONFIG_FILE)


.PHONY: deploy-postgres
deploy-postgres:
	@echo "Deploying postgres in Docker container"
	export POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) && \
		export POSTGRES_USER=$(POSTGRES_USER) && \
		export POSTGRES_DB=$(POSTGRES_DB) && \
		$(DOCKER_CMD) run -d --name reforged-labs-db -p 5432:5432 \
			-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
			-e POSTGRES_USER=$(POSTGRES_USER) \
			-e POSTGRES_DB=$(POSTGRES_DB) postgres

.PHONY: docker-build
docker-build:
	@echo "Building docker image"
	docker build -f Dockerfile --build-arg goproxy=$(GOPROXY) --build-arg ARCH=$(ARCH) --build-arg LDFLAGS="$(LDFLAGS)" . -t $(CONTROLLER_IMG_TAG)

.PHONY: docker-push
docker-push:
	@echo "Pushing docker image"
	docker push $(CONTROLLER_IMG_TAG)