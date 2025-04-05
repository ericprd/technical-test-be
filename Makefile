BINARY_NAME=./cmd/app/main.go
BINARY_MIGRATE_NAME=./cmd/gorm/main.go
OUTPUT_DIR=./bin/app

all: build

build:
	@echo "Building the Go project..."
	go build -o $(OUTPUT_DIR) $(BINARY_NAME)

run: build
	@echo "Building and running..."
	$(OUTPUT_DIR)

build-migrate:
	@echo "Building the Migrate..."
	go build -o $(OUTPUT_MIGRATE_DIR) $(BINARY_MIGRATE_NAME)

# migrate: build-migrate
# 	@echo "Building and running Migrate..."
# 	$(OUTPUT_MIGRATE_DIR)

migrate:
	go run $(BINARY_MIGRATE_NAME)

clean:
	@echo "Cleaning the build..."
	rm -f $(OUTPUT_DIR)
