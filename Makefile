SOURCE_FILES?=./...
GOLANGCI_VERSION=v1.49.0

export PATH := ./bin:$(PATH)
export GO111MODULE := on

.PHONY: fmt
fmt:
	@echo "==> Formatting all files..."
	find . -name '*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done