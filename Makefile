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

rel release: build
	mkdir -p release
	mkdir -p release/bin
	cp appspec.yml release
	cp -r scripts release
	cp bin/watcher  release/bin

.PHONY: all test clean build install

# vim:ft=make
