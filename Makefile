
scanner:
	@go build -o bin/scanner ./cmd/scanner
	@./bin/scanner

storageUnit:
	@go build -o bin/storageUnit ./cmd/storageUnit
	@./bin/storageUnit

aggregator:
	@go build -o bin/aggregator ./cmd/aggregator/
	@./bin/aggregator

api:
	@go build -o bin/api ./cmd/api/
	@./bin/api

test:
	@go test -count=1 -v ./...

