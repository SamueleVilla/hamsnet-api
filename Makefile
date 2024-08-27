include .env
POSTGRES_DSN=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable&connect_timeout=10
BINARY_NAME=bin/api/hamsnetapi

build:
	@echo "Building the executable..."
	go build -o ${BINARY_NAME} cmd/api/main.go
	@echo "Executable built at ${BINARY_NAME}"

run: build swagger
	@echo "Running the executable..."
	./${BINARY_NAME}

test:
	go test -v ./...

migrate_up:
	@echo "Migrating up..."
	 migrate -path=internal/database/migrations -database "${POSTGRES_DSN}" -verbose up

migrate_down:
	@echo "Migrating down..."
	migrate -path=internal/database/migrations -database "${POSTGRES_DSN}" -verbose down

migrate_create:
	@echo "Creating migration with name: $(name)"
	migrate create -ext sql -dir internal/database/migrations -seq $(name)

swagger:
	@echo "Generating swagger docs..."
	swag init -g ./internal/api/server.go