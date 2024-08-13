# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGENERATE=$(GOCMD) generate
BINARY_NAME=server

# Proto file
PROTO_FILES=$(PROTO_DIR)/*.proto

# Default target executed when no arguments are given to make
all: build

# Build the project
build: generate
	$(GOBUILD) -o $(BINARY_NAME) -v ./...

# Clean the build cache
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Generate gRPC code
generate:
	$(GOGENERATE)

# Install dependencies
deps:
	$(GOGET) -u google.golang.org/grpc
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go
	$(GOGET) -u github.com/cosmtrek/air

# Run the server
run: build
	./$(BINARY_NAME)

# Run the server with live-reloading using air
air:
	air

.PHONY: all build clean test generate deps run air