.DEFAULT_GOAL := help

GO       ?= go
GOFLAGS  ?=
PROJECT_NAME ?= memo
LOCAL_BIN = $(CURDIR)/bin

include bin-deps.mk

.PHONY: format
format: $(IMPREVISER_BIN) $(GOFUMPT_BIN) ## formats the code and also imports order
	@for f in $$(git status --porcelain | awk 'match($$1, "M|A|\\?"){print $$2}' | grep '.go$$'); do \
		$(IMPREVISER_BIN) -project-name $(PROJECT_NAME) $$f; \
		$(GOFUMPT_BIN) -l -w -extra $$f; \
	done

.PHONY: lint
lint: $(LINT_BIN) ## lints the code
	$(LINT_BIN) run --fix

.PHONY: install-githooks
install-githooks: ## installs all git hooks
	cp ./githooks/* .git/hooks/

.PHONY: wire
wire: $(WIRE_BIN) ## injects dependencies
	$(WIRE_BIN) cmd/polling_bot/wire.go 
	$(WIRE_BIN) cmd/rand_cmd/wire.go 

.PHONY: build
build: wire ## builds all commands
	$(GO) $(GOFLAG) build -o $(LOCAL_BIN)/polling_bot ./cmd/polling_bot
	$(GO) $(GOFLAG) build -o $(LOCAL_BIN)/rand_cmd ./cmd/rand_cmd

.PHONY: run
run: build ## builds and runs the polling bot
	$(LOCAL_BIN)/polling_bot

.PHONY: test
test: ## runs tests
	$(GO) $(GOFLAG) test ./...

.PHONY: db-up
db-up: ## initializes or migrates db
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

.PHONY: generate
generate: $(MOCKERY_BIN) wire ## generates code
	@PATH="$$PATH:$(LOCAL_BIN)" GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) generate ./...

.PHONY: help
help:
	@grep --no-filename -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
