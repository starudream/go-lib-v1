PROJECT ?= $(shell basename $(CURDIR))
MODULE  ?= $(shell go list -m)

GO      ?= GO111MODULE=on go
VERSION ?= $(shell git describe --tags 2>/dev/null || echo "dev")
BIDTIME ?= $(shell date +%FT%T%z)

BITTAGS := viper_logger
LDFLAGS := -s -w
LDFLAGS += -X "$(MODULE)/constant.VERSION=$(VERSION)"
LDFLAGS += -X "$(MODULE)/constant.BIDTIME=$(BIDTIME)"
LDFLAGS += -X "$(MODULE)/constant.PREFIX=app"

export APP_CONFIG_PATH = $(shell pwd)/config.yaml

.PHONY: tidy
tidy:
	@$(GO) mod tidy

.PHONY: example-simple
example-simple:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) run -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' $(MODULE)/example/simple

.PHONY: version-example-simple
version-example-simple:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) build -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' -o bin/example-simple $(MODULE)/example/simple
	go version -m bin/example-simple

.PHONY: lint
lint:
	@$(MAKE) tidy
	golangci-lint run --skip-dirs-use-default
