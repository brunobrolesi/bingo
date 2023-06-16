PACKAGES_PATH = $(shell go list -f '{{ .Dir }}' ./...)

.PHONY: all
all: require tidy fmt goimports vet staticcheck test

.PHONY: require
require:
	@type "goimports" > /dev/null 2>&1 \
		|| (echo 'goimports not found: to install it, run "go install golang.org/x/tools/cmd/goimports@latest"'; exit 1)
	@type "staticcheck" > /dev/null 2>&1 \
		|| (echo 'staticcheck not found: to install it, run "go install honnef.co/go/tools/cmd/staticcheck@latest"'; exit 1)

.PHONY: pre-commit-update
pre-commit-update:
	@echo "=> Executing pre-commit autoupdate"
	@pre-commit autoupdate

.PHONY: tidy
tidy:
	@echo "=> Executing go mod tidy"
	@go mod tidy

.PHONY: fmt
fmt:
	@echo "=> Executing go fmt"
	@go fmt ./...

.PHONY: goimports
goimports:
	@echo "=> Executing goimports"
	@goimports -w $(PACKAGES_PATH)

.PHONY: vet
vet:
	@echo "=> Executing go vet"
	@go vet ./...

.PHONY: staticcheck
staticcheck:
	@echo "=> Executing staticcheck"
	@staticcheck ./...

.PHONY: test
test:
	@echo "=> Running tests"
	@go test ./... -covermode=atomic -coverpkg=./... -count=1 -race -shuffle=on

.PHONY: test-cover
test-cover:
	@echo "=> Running tests and generating report"
	@go test ./... -covermode=atomic -coverprofile=/tmp/coverage.out -coverpkg=./... -count=1 -race -shuffle=on
	@go tool cover -html=/tmp/coverage.out

.PHONY: build-windows
build-windows:
	@echo "=> Building for Windows"
	@GOOS=windows GOARCH=amd64 go build -o bin/bingo.exe cmd/main.go

.PHONY: build-macos-intel
build-macos-intel:
	@echo "=> Building for MacOS Intel"
	@GOOS=darwin GOARCH=amd64 go build -o bin/bingo cmd/main.go

.PHONY: build-macos-arm
build-macos-arm:
	@echo "=> Building for MacOS ARM"
	@GOOS=darwin GOARCH=arm64 go build -o bin/bingo cmd/main.go

.PHONY: build-linux
build-linux:
	@echo "=> Building for Linux"
	@GOOS=linux GOARCH=amd64 go build -o bin/bingo cmd/main.go