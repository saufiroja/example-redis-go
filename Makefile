.PHONY: run-docker
build:
	@echo "building docker image..."
	@docker compose up -d

.PHONY: remove
remove:
	@echo "removing docker image..."
	@docker compose down -v

.PHONY: run
run:
	@echo "running app..."
	@go run main.go

.PHONY: test
test:
	@echo "running tests..."
	@go test -v ./test