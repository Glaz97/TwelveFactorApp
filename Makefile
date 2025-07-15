.PHONY: swag start

VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null)
BRANCH ?= $(shell git symbolic-ref -q --short HEAD 2>/dev/null)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)

GOARGS += -v
LDFLAGS += -s -w -X ${MODULE}/pkg/buildinfo.version=${VERSION} \
	-X ${MODULE}/pkg/buildinfo.commitHash=${COMMIT_HASH} \
	-X ${MODULE}/pkg/buildinfo.buildDate=${BUILD_DATE} \
	-X ${MODULE}/pkg/buildinfo.branch=${BRANCH}

.PHONY: build
build:
	go build ${GOARGS} -tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${OUTPUT} ${PACKAGE}

swag:
	swag fmt
	swag init -g internal/handler/handler.go --output pkg/docs

mongo:
	docker compose -f test/docker-compose.yaml -p twelvefactorapp up -d

mongo-down:
	docker compose -f test/docker-compose.yaml -p twelvefactorapp down

test: mongo
	go test -v -race ./...

start: mongo
	go run cmd/twelvefactorapp/main.go config  > configs/twelvefactorapp.yaml