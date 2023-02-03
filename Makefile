PROJECT ?= $(shell basename $(CURDIR))
MODULE  ?= $(shell go list -m)

GO      ?= GO111MODULE=on go
VERSION ?= $(shell git describe --tags 2>/dev/null || echo "dev")
BIDTIME ?= $(shell date +%FT%T%z)

BITTAGS :=
LDFLAGS := -s -w
LDFLAGS += -X "$(MODULE)/constant.VERSION=$(VERSION)"
LDFLAGS += -X "$(MODULE)/constant.BIDTIME=$(BIDTIME)"
LDFLAGS += -X "$(MODULE)/constant.NAME=starudream"
LDFLAGS += -X "$(MODULE)/constant.PREFIX=app"

.PHONY: bin
bin:
	@$(MAKE) bin-app bin-bot bin-server bin-simple

bin-%:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) build -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' -o bin/example-$* $(MODULE)/example/$*

.PHONY: example-simple
example-simple:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) run -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' $(MODULE)/example/simple

.PHONY: example-bot
example-bot:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) run -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' $(MODULE)/example/bot

.PHONY: example-server
example-server:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) run -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' $(MODULE)/example/server

.PHONY: tidy
tidy:
	@$(GO) mod tidy

.PHONY: lint
lint:
	@$(MAKE) tidy
	golangci-lint run --skip-dirs internal/viper
