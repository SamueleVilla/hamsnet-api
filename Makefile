build:
	go build -o bin/hasnetapi cmd/main.go

run: build
	./bin/hasnetapi

test:
	go test -v ./...