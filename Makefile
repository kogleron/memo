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
	${HOME}/go/bin/golangci-lint run

install-githooks:
	@echo "Installing githooks"
	cp ./githooks/* .git/hooks/

build:
	$(GO) $(GOFLAG) build -o ./bin/memo ./cmd/memo

run: build
	./bin/memo

init-db:
	$(GO) $(GOFLAG) run migrations/sqlite.go