.PHONY: build
build:
	go build

.PHONY: test
test:
	go test -v .

.PHONY: lint
lint:
	golangci-lint run ./... --config=.golangci.yml

.PHONY: lint-verbose
lint-verbose:
	golangci-lint run ./... -v --config=.golangci.yml

.PHONY: install-lint
install-lint:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
