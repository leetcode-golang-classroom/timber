build:
	@go build -o ./bin/timber

run: build
	@./bin/timber