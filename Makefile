ea Makefile
# giulio, 2015-10-21 09:30
#
# Simple make file that assumes a gb
# executable is available in path.
#
PROJECT = go-watcher
VERSION = 1.0.0

# gb 0.2.0 supports build flags
GOFLAGS ?= $(GOFLAGS:)

all: build test

build:
	@gb build $(GOFLAGS) ./...

test: build
	@gb test $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...

rel release:
	cp appspec.yml _build/prod/rel
	cp -r scripts _build/prod/rel
	cp config/config.sh _build/prod/rel

.PHONY: all test clean build install

# vim:ft=make
