build:
	@go build -o ./bin/timber

run: build
	@./bin/timber

test:
	@go test -v ./...