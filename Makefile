# Weather CLI Makefile

.PHONY: build clean test run demo install help

# Default target
help:
	@echo "Weather CLI - Beautiful Terminal Weather App"
	@echo ""
	@echo "Available targets:"
	@echo "  build     - Build the application"
	@echo "  clean     - Remove build artifacts"
	@echo "  test      - Run tests"
	@echo "  run       - Run the application"
	@echo "  demo      - Run demo with sample data"
	@echo "  install   - Install dependencies"
	@echo "  help      - Show this help message"

# Build the application
build:
	@echo "Building Weather CLI..."
	go build -o weather-cli$(EXT) .
	@echo "Build complete! Binary: weather-cli$(EXT)"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f weather-cli weather-cli.exe
	go clean

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Run the application
run:
	@echo "Running Weather CLI..."
	go run . --help

# Run demo
demo: build
	@echo "Running Weather CLI demo..."
ifeq ($(OS),Windows_NT)
	.\weather-cli.exe current "London"
	@echo ""
	.\weather-cli.exe forecast "Paris"
else
	./weather-cli current "London"
	@echo ""
	./weather-cli forecast "Paris"
endif

# Set extension based on OS
ifeq ($(OS),Windows_NT)
    EXT := .exe
else
    EXT :=
endif
