# Define variables
GO=go
BINARY_NAME=main
SOURCES=$(wildcard *.go)

# Default target
all: build

# Build target
build:
	$(GO) build -o $(BINARY_NAME) $(SOURCES)

# Run target
run:
	$(GO) run $(SOURCES)

# Clean target
clean:
	$(RM) $(BINARY_NAME)
