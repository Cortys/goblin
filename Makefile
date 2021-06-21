# variable definitions
NAME := goblin
DESC := dumps a Go AST to JSON
PREFIX ?= usr/local
VERSION := $(shell git describe --tags --always --dirty)
GOVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDDATE := $(shell date -u +"%B %d, %Y")
BUILDER := $(shell echo "`git config user.name` <`git config user.email`>")
PKG_RELEASE ?= 1
PROJECT_URL := "https://github.com/Cortys/$(NAME)"
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildTime=$(BUILDTIME)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(GOVERSION)'

# development tasks
fmt:
	go fmt -x ./...

test: fmt
	go test -v $$(go list ./... | grep -v /vendor/ | grep -v /cmd/ | grep -v /fixtures/)

PACKAGES := $(shell find ./* -type d | grep -v vendor)

coverage:
	@go test -coverprofile=coverage.txt -covermode=atomic
	@-go tool cover -html=coverage.txt -o cover.html

benchmark:
	@echo "Running tests..."
	@go test -bench=. $$(go list ./... | grep -v /vendor/ | grep -v /cmd/)

CMD_SOURCES := $(shell find cmd -name main.go)
TARGETS := $(patsubst cmd/%/main.go,%,$(CMD_SOURCES))

%: cmd/%/main.go
	go build -ldflags "$(LDFLAGS)" -o $@ $<

%-${TRAVIS_TAG}-${TARGET}.tar.gz: %
	tar czf $@ $<

.PHONY: fmt