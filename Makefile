.DEFAULT_GOAL := help

GO       ?= go
GOFLAGS  ?=
PROJECT_NAME ?= memo
LOCAL_BIN = $(CURDIR)/bin

.PHONY: install
install: ## installs dependencies
	@echo "Install required programs"
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install golang.org/x/tools/cmd/goimports@latest
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install mvdan.cc/gofumpt@latest
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) get -v github.com/incu6us/goimports-reviser
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/google/wire/cmd/wire@latest

.PHONY: format
format: ## formats the code and also imports order
	@echo "Formatting..."
	$(LOCAL_BIN)/gofumpt -l -w -extra .
	@echo "Formatting imports..."
	@for f in $$(find . -name '*.go'); do \
		$(LOCAL_BIN)/goimports-reviser -file-path $$f -project-name $(PROJECT_NAME); \
	done

.PHONY: lint
lint: ## lints the code
	@echo "Linting"
	$(LOCAL_BIN)/golangci-lint run --fix

.PHONY: install-githooks
install-githooks: ## installs all git hooks
	@echo "Installing githooks"
	cp ./githooks/* .git/hooks/

.PHONY: wire
wire: ## injects dependencies
	$(LOCAL_BIN)/wire cmd/polling_bot/wire.go 
	$(LOCAL_BIN)/wire cmd/rand_cmd/wire.go 

.PHONY: build
build: wire ## builds all commands
	$(GO) $(GOFLAG) build -o $(LOCAL_BIN)/polling_bot ./cmd/polling_bot
	$(GO) $(GOFLAG) build -o $(LOCAL_BIN)/rand_cmd ./cmd/rand_cmd

.PHONY: run
run: build ## builds and runs the polling bot
	$(LOCAL_BIN)/polling_bot

.PHONY: test
test: ## runs tests
	@echo "Testing"
	$(GO) $(GOFLAG) test ./...

.PHONY: init-db
init-db: ## initializes db
	$(GO) $(GOFLAG) run migrations/sqlite.go

.PHONY: install-cron
install-cron: ## installs cron tasks
	PWD=$$(pwd); sed "s|\$$PROJECT_PATH|${PWD}|" cron/com.memo.PollingBot.plist >~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl bootout gui/$$UID ~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl bootstrap gui/$$UID ~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl kickstart gui/$$UID/com.memo.PollingBot
	PWD=$$(pwd); sed "s|\$$PROJECT_PATH|${PWD}|" cron/com.memo.RandCmd.plist >~/Library/LaunchAgents/com.memo.RandCmd.plist
	launchctl bootout gui/$$UID ~/Library/LaunchAgents/com.memo.RandCmd.plist
	launchctl bootstrap gui/$$UID ~/Library/LaunchAgents/com.memo.RandCmd.plist
	launchctl kickstart gui/$$UID/com.memo.RandCmd

.PHONY: help
help:
	@grep --no-filename -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
