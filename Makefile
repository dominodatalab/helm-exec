all: vet lint test

.PHONY: vet
vet:
	@go vet ./...

.PHONY: lint
lint:
	@golint ./...

.PHONY: test
test:
	@go test -cover -race ./...
