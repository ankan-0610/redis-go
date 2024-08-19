run: build
	@./bin/redis-go --listenAddr :3000

build:
	@go build -o bin/redis-go .