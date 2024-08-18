build:
	go build -o bin/api/hasnetapi cmd/api/main.go

run: build
	./bin/api/hasnetapi

test:
	go test -v ./...