CONTAINER_NAME=go_workspace
DOCKER_COMPOSE=docker-compose
APP_NAME=photo-extractor
BIN_DIR=bin

.PHONY: up down shell run tidy init build-linux build-windows build-all clean

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

shell:
	docker exec -it $(CONTAINER_NAME) sh

run:
	docker exec -it $(CONTAINER_NAME) go run ./cmd/extractor

tidy:
	docker exec -it $(CONTAINER_NAME) go mod tidy

init:
	docker exec -it $(CONTAINER_NAME) go mod init $(APP_NAME)

build-linux:
	docker exec -it $(CONTAINER_NAME) mkdir -p $(BIN_DIR)
	docker exec -e GOOS=linux -e GOARCH=amd64 $(CONTAINER_NAME) go build -o $(BIN_DIR)/$(APP_NAME)-linux ./cmd/extractor
	docker exec $(CONTAINER_NAME) chmod +x $(BIN_DIR)/$(APP_NAME)-linux

build-windows:
	docker exec -it $(CONTAINER_NAME) mkdir -p $(BIN_DIR)
	docker exec -e GOOS=windows -e GOARCH=amd64 $(CONTAINER_NAME) go build -o $(BIN_DIR)/$(APP_NAME).exe ./cmd/extractor

build-all: build-linux build-windows