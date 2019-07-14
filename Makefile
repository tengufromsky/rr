.PHONY: build
build:
	go build

.PHONY: test
test:
	go test -v .

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: install-lint
install-lint:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
