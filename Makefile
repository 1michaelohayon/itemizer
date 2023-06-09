
scanner:
	@go build -o bin/scanner ./cmd/scanner
	@./bin/scanner

storageUnit:
	@go build -o bin/storageUnit ./cmd/storageUnit
	@./bin/storageUnit

aggregator:
	@go build -o bin/aggregator ./cmd/aggregator/
	@./bin/aggregator
	
