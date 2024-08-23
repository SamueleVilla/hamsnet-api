include .env

build:
	go build -o bin/api/hasnetapi cmd/api/main.go

run: build swagger
	./bin/api/hasnetapi

test:
	go test -v ./...

migrate_up:
	 migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migrate_down:
	migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

migrate_create:
	migrate create -ext sql -dir internal/database/migrations -seq $(name)

swagger:
	swag init -g ./internal/api/server.go