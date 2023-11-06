.PHONY: test

test: start-local
	go test -cover -race -v ./...

build:
	@docker compose -f docker-compose.yml build 

start-local:
	@docker compose -f docker-compose.yml up -d

stop-local:
	@docker compose -f docker-compose.yml down