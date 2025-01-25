.PHONY: help all build test run dev clean install uninstall format lint

# Variables
BINARY_NAME = gocli
BINARY_PATH = ./bin/$(BINARY_NAME)
INSTALL_PATH = /usr/local/bin/$(BINARY_NAME)

help: # Display help
    @echo "Available targets:"
    @awk 'BEGIN {FS = ":.*?#"} /^[a-zA-Z_-]+:.*?#/ {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Run tests and build the project
all: test build

build: # Build the project and place the binary in the bin directory
    go build -o $(BINARY_PATH) ./cmd/$(BINARY_NAME)

test: # Run tests
    go test ./...

run: # Run the CLI application
    $(BINARY_PATH)

dev: # Run the application in development mode
    go run ./cmd/$(BINARY_NAME)

clean: # Clean up the binary
    rm -f $(BINARY_PATH)

install: build # Install the binary to a system-wide location
    cp $(BINARY_PATH) $(INSTALL_PATH)

uninstall: # Uninstall the binary from the system-wide location
    rm -f $(INSTALL_PATH)

format: # Format the code
    go fmt ./...

lint: # Lint the code
    go vet ./...
