default: test

ACTIONLINT_VERSION := v1.7.7

acctest: fmt
	cd ./sdk/ && ACCTEST=1 go test -count=1 -race -v ./... -run '^TestAcc' && cd ../

fmt: tools
	@echo "==> Fixing source code with gofmt..."
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

imports: tools
	@echo "==> Fixing source code with goimports..."
	goimports -w ./data-plane
	goimports -w ./microsoft-graph
	goimports -w ./resource-manager
	goimports -w ./sdk

prepare:
	@echo "==> Preparing the repository to be updated.."
	@echo "==> 1. Resource Manager - Removing all existing generated files"
	find ./resource-manager/ -maxdepth 1 -mindepth 1 -type d -exec rm -rf '{}' \;
	@echo "==> 2. Resource Manager - Re-creating the directory structure in preparation"
	mkdir -p ./resource-manager/
	# TODO: enable Microsoft Graph support prior to enabling generation of that SDK
	#@echo "==> 3. Microsoft Graph - Removing all existing generated files"
	#find ./microsoft-graph/ -maxdepth 1 -mindepth 1 -type d -exec rm -rf '{}' \;
	#@echo "==> 4. Microsoft Graph - Re-creating the directory structure in preparation"
	#mkdir -p ./microsoft-graph/

test: fmt
	cd ./sdk/ && go test -short -v ./... && cd ../
	cd ./resource-manager/ && go test -v ./... && cd ../
	cd ./microsoft-graph/ && go test -v ./... && cd ../

test-ci-sdk:
	cd ./sdk/ && go test -short -v ./...

test-ci-resource-manager:
	cd ./resource-manager/ && go test -v ./...

test-ci-microsoft-graph:
	cd ./microsoft-graph/ && go test -v ./...

test-ci-data-plane:
	cd ./data-plane/ && go test -v ./...

test-ci: test-ci-sdk test-ci-resource-manager test-ci-microsoft-graph test-ci-data-plane

tools:
	@echo "==> installing required tooling..."
	go install golang.org/x/tools/cmd/goimports@v0.30.0
	go install github.com/rhysd/actionlint/cmd/actionlint@$(ACTIONLINT_VERSION)
	@if ! command -v shellcheck >/dev/null 2>&1; then \
		if command -v brew >/dev/null 2>&1; then \
			brew install shellcheck; \
		else \
			echo "Note: install shellcheck (e.g. 'brew install shellcheck' or 'apt-get install shellcheck') for full local linting."; \
		fi; \
	fi
	@if ! command -v zizmor >/dev/null 2>&1; then \
		if command -v brew >/dev/null 2>&1; then \
			brew install zizmor; \
		elif command -v pipx >/dev/null 2>&1; then \
			pipx install zizmor; \
		else \
			echo "Note: install zizmor (e.g. 'brew install zizmor' or 'pipx install zizmor') for workflow security linting."; \
		fi; \
	fi

validate-workflows:
	@echo "==> Running workflow validation..."
	@if command -v actionlint >/dev/null 2>&1; then \
		actionlint; \
	else \
		echo "actionlint not found. Run 'make tools' to install actionlint@$(ACTIONLINT_VERSION)."; \
		exit 1; \
	fi
	@if command -v shellcheck >/dev/null 2>&1; then \
		shellcheck -S error scripts/*.sh; \
	else \
		echo "shellcheck not found. Install it (e.g. 'brew install shellcheck') to lint scripts/*.sh."; \
		exit 1; \
	fi
	@if command -v zizmor >/dev/null 2>&1; then \
		zizmor .github/workflows/; \
	else \
		echo "Note: optional 'zizmor' not installed; skipping zizmor check."; \
	fi
	@./scripts/test-workflows.sh

# Optional local runner emulation via nektos/act (ubuntu-latest jobs only).
# Usage: make test-workflow-act WORKFLOW=.github/workflows/pr-validate-workflows.yaml [JOB=validate] [EVENT=pull_request] [ACT_ARGS="--dryrun"]
test-workflow-act:
	@if [ -z "$(WORKFLOW)" ]; then \
		echo "Error: WORKFLOW is required."; \
		echo "Usage: make test-workflow-act WORKFLOW=.github/workflows/pr-validate-workflows.yaml [JOB=validate] [EVENT=pull_request] [ACT_ARGS='--dryrun']"; \
		echo "Note: custom-linux-xl, Protected Environments, and OIDC jobs cannot be emulated locally by act."; \
		exit 1; \
	fi
	@command -v act >/dev/null 2>&1 || { echo "Error: 'act' not found. Install via 'brew install act'."; exit 1; }
	@docker info >/dev/null 2>&1 || { echo "Error: Docker daemon is not running."; exit 1; }
	@ARCH_FLAG=""; \
	if [ "$$(uname -m)" = "arm64" ]; then ARCH_FLAG="--container-architecture linux/amd64"; fi; \
	JOB_FLAG=""; \
	if [ -n "$(JOB)" ]; then JOB_FLAG="-j $(JOB)"; fi; \
	EVENT_NAME="$(or $(EVENT),pull_request)"; \
	act "$$EVENT_NAME" -W "$(WORKFLOW)" $$JOB_FLAG $$ARCH_FLAG -P ubuntu-latest=catthehacker/ubuntu:act-latest $(ACT_ARGS)

.PHONY: fmt imports prepare test test-ci test-ci-sdk test-ci-resource-manager test-ci-microsoft-graph test-ci-data-plane tools validate-workflows test-workflow-act
