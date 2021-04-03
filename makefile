APP=safe
PROJECT=github.com/nobe4/${APP}

# YYYY.MM.Count
VERSION?=2021.04.2
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

default: | build

version:
	@echo -n ${VERSION}

app-name:
	@echo -n ${APP}

build:
	goreleaser --snapshot --rm-dist

# Used for manually releasing, normally running in github actions.
release:
ifneq ($(shell git symbolic-ref --short HEAD),master)
	$(error Not on master branch)
endif

	goreleaser --rm-dist

lint:
	golangci-lint run

test:
	go test ./...

bump:
	./scripts/bump.sh
	git add makefile
	git commit -m "Bump version" --edit

tag:
	./scripts/tag.sh

.PHONY: bump tag lint test app-name version
