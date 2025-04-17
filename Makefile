APP_NAME := $(notdir $(CURDIR))
BUILD_DIR := bin
CMD_DIR := cmd/$(APP_NAME)

.PHONY: all build clean run test

all: build

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME).out $(CMD_DIR)/main.go

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

run: build
	@echo "Running $(APP_NAME)..."
	./$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "Running tests..."
	go test ./...
