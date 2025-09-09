# httpcodes - HTTP Status Code Lookup Tool
# Data sourced from: https://github.com/MattIPv4/status-codes

BINARY_NAME=httpcodes
BUILD_DIR=build
GO_FILES=main.go

.PHONY: all build clean install test run help

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_FILES)
	@echo "Built $(BUILD_DIR)/$(BINARY_NAME)"

install: build
	@echo "Installing $(BINARY_NAME)..."
	@go install
	@echo "Installed $(BINARY_NAME) to $$GOPATH/bin or $$HOME/go/bin"

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean

run:
	@go run $(GO_FILES) $(ARGS)

test: build
	@echo "Testing with common HTTP status codes..."
	@echo "Testing 200:"
	@./$(BUILD_DIR)/$(BINARY_NAME) 200
	@echo "Testing 404:"
	@./$(BUILD_DIR)/$(BINARY_NAME) 404
	@echo "Testing 500:"
	@./$(BUILD_DIR)/$(BINARY_NAME) 500

# Show help
help:
	@echo "httpcodes - HTTP Status Code Lookup Tool"
	@echo ""
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  install  - Install the binary to GOPATH/bin"
	@echo "  clean    - Clean build artifacts"
	@echo "  run      - Run the application (use ARGS=<code>)"
	@echo "  test     - Test with common status codes"
	@echo "  help     - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make run ARGS=200"
	@echo "  make install"
	@echo "  make test"
