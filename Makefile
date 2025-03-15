GO_CMD ?= "/Volumes/veronica/git-cuongpiger/go-env/go-1.24.1/go1.24.1/bin/go"


.PHONY: run-api-service
run-api-service:
	@echo "Running API service at localhost"
	$(GO_CMD) run cmd/api-service/main.go
