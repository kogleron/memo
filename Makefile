GO       ?= go
GOFLAGS  ?=
PROJECT_NAME ?= memo

install:
	@echo "Install required programs"
	$(GO) $(GOFLAG) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) $(GOFLAG) install golang.org/x/tools/cmd/goimports@latest
	$(GO) $(GOFLAG) install mvdan.cc/gofumpt@latest
	$(GO) $(GOFLAG) get -v github.com/incu6us/goimports-reviser

format:
	@echo "Formatting..."
	${HOME}/go/bin/gofumpt -l -w -extra .
	@echo "Formatting imports..."
	@for f in $$(find . -name '*.go'); do \
		${HOME}/go/bin/goimports-reviser -file-path $$f -project-name $(PROJECT_NAME); \
	done

lint:
	@echo "Linting"
	${HOME}/go/bin/golangci-lint run --fix

install-githooks:
	@echo "Installing githooks"
	cp ./githooks/* .git/hooks/

build:
	$(GO) $(GOFLAG) build -o ./bin/polling_bot ./cmd/polling_bot

run: build
	./bin/memo

test:
	@echo "Testing"
	$(GO) $(GOFLAG) test ./...

init-db:
	$(GO) $(GOFLAG) run migrations/sqlite.go

install-cron:
	PWD=$$(pwd); sed "s|\$$PROJECT_PATH|${PWD}|" cron/com.memo.PollingBot.plist >~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl bootout gui/$$UID ~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl bootstrap gui/$$UID ~/Library/LaunchAgents/com.memo.PollingBot.plist
	launchctl kickstart gui/$$UID/com.memo.PollingBot
