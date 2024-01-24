default: test

acctest: fmt
	cd ./sdk/ && ACCTEST=1 go test -count=1 -race -v ./... -run '^TestAcc' && cd ../

fmt: tools
	@echo "==> Fixing source code with gofmt..."
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

imports: tools
	@echo "==> Fixing source code with goimports..."
	goimports -w ./microsoft-graph
	goimports -w ./resource-manager
	goimports -w ./sdk

prepare:
	@echo "==> Preparing the repository to be updated.."
	@echo "==> 1. Resource Manager - Removing all existing generated files"
	rm -rf ./resource-manager/
	@echo "==> 2. Resource Manager - Re-creating the directory structure in preparation"
	mkdir -p ./resource-manager/
	#@echo "==> 3. Microsoft Graph - Removing all existing generated files"
	#rm -rf ./microsoft-graph/
	#@echo "==> 4. Microsoft Graph - Re-creating the directory structure in preparation"
	#mkdir -p ./microsoft-graph/

test: fmt
	cd ./sdk/ && go test -short -v ./... && cd ../
	cd ./resource-manager/ && go test -v ./... && cd ../
	cd ./microsoft-graph/ && go test -v ./... && cd ../

tools:
	@echo "==> installing required tooling..."
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: fmt imports prepare test tools
