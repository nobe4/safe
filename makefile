APP=safe
PROJECT=github.com/nobe4/${APP}

GO?=go
GOOS?=darwin
GOARCH?=amd64

# YYYY.MM.Count
RELEASE?=2021.04.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

BUILD_PATH=./bin/${GOOS}-${GOARCH}/${APP}
MAIN_PATH=${PROJECT}/cmd/${APP}

default: | build

.PHONY: version
version:
	@echo -n ${RELEASE}

.PHONY: app-name
app-name:
	@echo -n ${APP}

build:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	go build -a \
		-ldflags="-s -w \
			-X 'main.Version=${RELEASE}' \
			-X 'main.Commit=${COMMIT}' \
			-X 'main.Build=${BUILD_TIME}' "\
		-o ${BUILD_PATH} ${MAIN_PATH}

lint:
	golangci-lint run

test:
	go test ./...
