#
# @author Jose Nidhin
#

VERSION := $(shell git describe --always --long --dirty)
PROJECT_NAME := $(shell basename "$(PWD)")

GO_SRC_FILES := $(shell find . -type f -name '*.go')
GO_SRC_SERVER := $(shell ls cmd/server/*.go)

GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

GOSTRINGER_VERSION := v0.26.0
GOSTRINGER_EXE := $(GOBIN)/stringer

GOLANGCI_VERSION := v1.61.0
GOLANGCI_EXE := $(GOBIN)/golangci-lint

GITLEAKS_VERSION := v8@latest
GITLEAKS_EXE := $(GOBIN)/gitleaks

CONTAINER_VERSION := latest
GO_ENVFLAGS := CGO_ENABLED=0

LDFLAGS := -X git.dtone.com/nexus/connexus/buildinfo.AppVersion=$(VERSION)
LDFLAGS += -X git.dtone.com/nexus/connexus/buildinfo.AppName=$(PROJECT_NAME)

GO_TAG_PRIVATE_ROUTE_HANDLER := private_route_handler

CMD_ROUTE_PATH := cmd/route/main.go

all: generate tidy vet fmt simplify lint test gitleaks clean build

.PHONY: help
help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%s # %s\n", $$1, $$2}' | column -ts '#'

.PHONY: vet
vet: ## Runs the go vet command against the source code to check for issues.
	go vet -tags "$(GO_TAG_PRIVATE_ROUTE_HANDLER)" ./... && \
	go vet ./...

.PHONY: fmt
fmt: ## Uses go fmt command to standardize code format.
	gofmt -l -w $(GO_SRC_FILES)

.PHONY: simplify
simplify: ## Uses go fmt command with simplify option.
	gofmt -s -l -w $(GO_SRC_FILES)

.PHONY: tidy
tidy: ## Update the go.mod to match the source code.
	go mod tidy

.PHONY: generate
generate: ## Execute all the go generate directives within the source code.
	test -e $(GOSTRINGER_EXE) || $(info Installing go stringer tool.) \
	go install golang.org/x/tools/cmd/stringer@$(GOSTRINGER_VERSION)
	go generate ./...

.PHONY: test
test: ## Run all the test cases with race and cover option.
	go test -race -cover -tags "$(GO_TAG_PRIVATE_ROUTE_HANDLER)" ./...

.PHONY: clean
clean: ## Delete all generated artifacts.
	rm -f $(PROJECT_NAME) cover.out coverage.html

.PHONY: build
build: ## Build the private server executable of this source code.
	go build -v -ldflags "$(LDFLAGS)" -o $(PROJECT_NAME) -tags "$(GO_TAG_PRIVATE_ROUTE_HANDLER)" $(GO_SRC_SERVER)

.PHONY: build-public
build-public: ## Build the public server executable of this source code.
	go build -v -ldflags "$(LDFLAGS)" -o $(PROJECT_NAME) $(GO_SRC_SERVER)

.PHONY: coverage
coverage: ## Generate an html coverage report.
	go test -tags "$(GO_TAG_PRIVATE_ROUTE_HANDLER)" -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o coverage.html

.PHONY: release-build
release-build: ## Build the private server executable of this source code meant for production usage.
	$(GO_ENVFLAGS) go build -v -ldflags "$(LDFLAGS)" -o $(PROJECT_NAME) -tags "$(GO_TAG_PRIVATE_ROUTE_HANDLER)" $(GO_SRC_SERVER)

.PHONY: release-build-public
release-build-public: ## Build the public server executable of this source code meant for production usage.
	$(GO_ENVFLAGS) go build -v -ldflags "$(LDFLAGS)" -o $(PROJECT_NAME) $(GO_SRC_SERVER)

.PHONY: container-clean-build
container-clean-build: ## Build the container image for this project without using cache.
	docker image build --no-cache -t $(PROJECT_NAME):$(CONTAINER_VERSION) .

.PHONY: container-build
container-build: ## Build the container image for this project.
	docker image build -t $(PROJECT_NAME):$(CONTAINER_VERSION) .

.PHONY: lint-install
lint-install: ## Install the lint tool.
	test -e $(GOLANGCI_EXE) || $(info Installing golangci-lint.) \
	curl --silent \
	--show-error \
	--fail \
	--location \
	https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) $(GOLANGCI_VERSION)
	$(GOLANGCI_EXE) --version

.PHONY: lint-uninstall
lint-uninstall: ## Uninstall the lint tool.
	rm $(GOLANGCI_EXE)

.PHONY: lint-tool-version
lint-tool-version: ## Print the version of the lint tool.
	$(GOLANGCI_EXE) --version

.PHONY: lint
lint: lint-install fmt lint-tool-version lint-private-routes lint-public-routes ## Run lint tool against the source code. (Recommended).

.PHONY: lint-private-routes
lint-private-routes: ## Run lint tool against source code with --build-tag set to compile private routes. (Recommended: lint).
	$(GOLANGCI_EXE) run -v --build-tags "$(GO_TAG_PRIVATE_ROUTE_HANDLER)"

.PHONY: lint-public-routes
lint-public-routes: ## Run lint tool against source code with --build-tag set to compile public routes. (Recommended: lint).
	$(GOLANGCI_EXE) run -v

.PHONY: lint-fix
lint-fix: lint-install fmt ## Run the lint tool against the source code with auto-fix enabled.
	$(GOLANGCI_EXE) --version
	$(GOLANGCI_EXE) run -v --fix --build-tags "$(GO_TAG_PRIVATE_ROUTE_HANDLER)" && \
  $(GOLANGCI_EXE) run -v --fix

.PHONY: print-routes
print-routes: ## Print all the HTTP routes supported by this source code.
	go run -ldflags "$(LDFLAGS)" -tags "$(GO_TAG_PRIVATE_ROUTE_HANDLER)" $(CMD_ROUTE_PATH)

.PHONY: gitleaks-install
gitleaks-install: ## Install the gitleaks tool.
	test -e $(GITLEAKS_EXE) || $(info Installing gitleaks tool.) \
  	go install github.com/zricethezav/gitleaks/$(GITLEAKS_VERSION)
	$(GITLEAKS_EXE) --version

.PHONY: gitleaks-uninstall
gitleaks-uninstall: ## Uninstall the gitleaks tool.
	rm $(GITLEAKS_EXE)

.PHONY: gitleaks-tool-version
gitleaks-tool-version: ## Print the version of the gitleaks tool.
	$(GITLEAKS_EXE) --version

.PHONY: gitleaks
gitleaks: gitleaks-install gitleaks-tool-version gitleaks-scan ## Run gitleaks tool against the source code. (Recommended).

.PHONY: gitleaks-scan
gitleaks-scan: ## Gitleaks scan secrets at unstaged, staged and previous commit.
	$(GITLEAKS_EXE) git --pre-commit --config gitleaks.toml -v
	$(GITLEAKS_EXE) git --staged --config gitleaks.toml -v
	$(GITLEAKS_EXE) git --config gitleaks.toml -v

.PHONY: setup-pre-commit-hook
setup-pre-commit-hook: ## Setup git pre-commit hook.
	cd ./.git/hooks && ln -s -f ../../git_hooks/pre-commit