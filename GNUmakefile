default: test

fmt: tools
	@echo "==> Fixing source code with gofmt..."
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

imports: tools
	@echo "==> Fixing source code with goimports..."
	goimports -w ./resource-manager

prepare:
	@echo "==> Preparing the repository to be updated.."
	@echo "==> 1. Removing all existing generated files"
	rm -rf ./resource-manager/
	@echo "==> 2. Re-creating the directory structure in preparation"
	mkdir -p ./resource-manager/

test: fmt
	go test -v ./resource-manager/...
	go test -v ./sdk/...

tools:
	@echo "==> installing required tooling..."
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: fmt imports prepare test tools
