.PHONY: test

test:
	go test -cover -race -v ./...

# lint:
# 	golangci-lint run -p bugs -p error  

start-local:
	@docker compose -f docker-compose.yml up -d

stop-local:
	@docker compose -f docker-compose.yml down