
WIRE_BIN=$(LOCAL_BIN)/wire
$(WIRE_BIN):
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/google/wire/cmd/wire@v0.5.0

GOFUMPT_BIN=$(LOCAL_BIN)/gofumpt
$(GOFUMPT_BIN):
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install mvdan.cc/gofumpt@v0.4.0

IMPREVISER_BIN=$(LOCAL_BIN)/goimports-reviser
$(IMPREVISER_BIN):
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/incu6us/goimports-reviser/v3@v3.3.0

LINT_BIN=$(LOCAL_BIN)/golangci-lint
$(LINT_BIN):
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2

MOCKERY_BIN=$(LOCAL_BIN)/mockery
$(MOCKERY_BIN):
	GOBIN=$(LOCAL_BIN) $(GO) $(GOFLAG) install github.com/vektra/mockery/v2@v2.20.2
