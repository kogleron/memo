.DEFAULT_GOAL := help

GO       ?= go
GOFLAGS  ?=
PROJECT_NAME ?= memo
LOCAL_BIN = $(CURDIR)/bin

.PHONY: install
install: ## installs dependencies
	@echo "Install required programs"
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install golang.org/x/tools/cmd/goimports@latest
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install mvdan.cc/gofumpt@latest
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install -v github.com/incu6us/goimports-reviser/v3@v3.3.1
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/google/wire/cmd/wire@v0.5.0

.PHONY: format
format: ## formats the code and also imports order
	@echo "Formatting..."
	@for f in $$(git status --porcelain | awk 'match($$1, "M|A|\\?"){print $$2}' | grep '.go$$'); do \
		$(IMPREVISER_BIN) -project-name $(PROJECT_NAME) $$f; \
		$(GOFUMPT_BIN) -l -w -extra $$f; \
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

.PHONY: cron-install
cron-install: ## installs cron tasks
	PWD=$$(pwd); sed "s|\$$PROJECT_PATH|${PWD}|" cron/com.memo.PollingBot.plist >~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl bootout gui/$$UID ~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl bootstrap gui/$$UID ~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl kickstart gui/$$UID/com.memo.PollingBot
	PWD=$$(pwd); sed "s|\$$PROJECT_PATH|${PWD}|" cron/com.memo.RandCmd.plist >~/Library/LaunchAgents/com.memo.RandCmd.plist
	launchctl bootout gui/$$UID ~/Library/LaunchAgents/com.memo.RandCmd.plist
	launchctl bootstrap gui/$$UID ~/Library/LaunchAgents/com.memo.RandCmd.plist
	launchctl kickstart gui/$$UID/com.memo.RandCmd

.PHONY: cron-uninstall
cron-uninstall: ## uninstalls cron tasks
	launchctl bootout gui/$$UID ~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl bootout gui/$$UID ~/Library/LaunchAgents/com.memo.RandCmd.plist

.PHONY: help
help:
	@grep --no-filename -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
