.PHONY: lint
lint:
	golangci-lint run

.PHONY: generate
generate:
	./scripts/gen_proto
	go generate ./...

.PHONY: test
test:
	go test ./... -cover