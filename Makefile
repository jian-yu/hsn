#!/usr/bin/make -f

VERSION := $(shell echo $(shell git rev-parse HEAD) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true

export GO111MODULE = on

# process build tags

build_tags = netgo

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=hsn \
		  -X github.com/cosmos/cosmos-sdk/version.ServerName=hsnd \
		  -X github.com/cosmos/cosmos-sdk/version.ClientName=hsncli \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'


all: install

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/hsnd.exe ./cmd/hsnd
	go build -mod=readonly $(BUILD_FLAGS) -o build/hsncli.exe ./cmd/hsncli
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/hsnd ./cmd/hsnd
	go build -mod=readonly $(BUILD_FLAGS) -o build/hsncli ./cmd/hsncli
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hsnd
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hsncli

########################################
### Tools & dependencies

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

clean:
	rm -rf build/

.PHONY: all build-linux install clean build