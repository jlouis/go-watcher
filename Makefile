#
# Makefile
# giulio, 2015-10-21 09:30
#
# Simple make file that assumes a gb
# executable is available in path.
.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

all: install test

build:
	@gb build $(GOFLAGS) ./...

test: install
	@gb test $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...

# vim:ft=make
#
