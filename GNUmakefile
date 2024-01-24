default: test

acctest: fmt
	ACCTEST=1 go test -count=1 -race -v ./sdk/... -run '^TestAcc'

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
	cd ./sdk/ && go test -short -v ./... && cd ../
	#cd ./resource-manager/ && go test -v ./... && cd ../

tools:
	@echo "==> installing required tooling..."
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: fmt imports prepare test tools
