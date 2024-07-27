default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done


arch?=$(shell go env GOARCH)
os?=$(shell go env GOOS)
ARCH=$(arch)
OS=$(os)
CGO=1
BIN_FOLDER=bin
PROJECT_FOLDER=.
CMD_FOLDER=./cmd


VERSION:=$(shell cat VERSION)
HASH = $(shell git rev-parse --short HEAD)
DIRTY = $(shell bash -c 'if [ -n "$$(git status --porcelain --untracked-files=no)" ]; then echo -dirty; fi')
COMMIT ?= $(HASH)$(DIRTY)
Built:=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS:=-X main.Version=$(VERSION) \
        -X main.Commit=$(COMMIT) \
        -X main.Built=$(Built)



dev-dep: localredis localnats localmongodb app-dep  # Start all services for development
dev-build: readapi writeapi	_scraper
dev-run: dev-dep dev-build run


k8s: redis-helm-install  searchengine-helm-install
k8su: searchengine-helm-uninstall


.PHONY: localredis
localredis: # Start a redis-server
	screen -S redis -dm redis-stack-server

localnats:
	screen -S nats-server -dm nats-server -js

localmongodb: # Start a mongodb
	mongod --config /opt/homebrew/etc/mongod.conf --fork

protoc: # Generate client and server code
	 protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        protos/searchreq/searchreq.proto
	protoc --go_out=. --go_opt=paths=source_relative \
            --go-grpc_out=. --go-grpc_opt=paths=source_relative \
            protos/backendreq/backendreq.proto
	protoc --go_out=. --go_opt=paths=source_relative \
              --go-grpc_out=. --go-grpc_opt=paths=source_relative \
              protos/event/event.proto

readapi: # Starts reada pi.
	@echo "  >  Building binary for ${OS}-${ARCH}"
	CGO_ENABLED=${CGO} GOOS=${OS} GOARCH=${ARCH} go build -C ${PROJECT_FOLDER} -ldflags="${LDFLAGS}" \
			-o "${BIN_FOLDER}/readapi" "${CMD_FOLDER}/api/read"


writeapi: # Starts write api.
	@echo "  >  Building binary for ${OS}-${ARCH}"
	CGO_ENABLED=${CGO} GOOS=${OS} GOARCH=${ARCH} go build -C ${PROJECT_FOLDER} -ldflags="${LDFLAGS}" \
			-o "${BIN_FOLDER}/writeapi" "${CMD_FOLDER}/api/write"


_scraper: # Starts scraper.
	@echo "  >  Building binary for ${OS}-${ARCH}"
	CGO_ENABLED=${CGO} GOOS=${OS} GOARCH=${ARCH} go build -C ${PROJECT_FOLDER} -ldflags="${LDFLAGS}" \
			-o "${BIN_FOLDER}/scraper" "${CMD_FOLDER}/scraper"

app-dep: # Starts app.
	cd web && npm install


run: #Run whole microservices.
	./bin/readapi --config configs/read/config.yaml & 2>/dev/null
	./bin/writeapi --config configs/write/config.yaml & 2>/dev/null
	./bin/scraper --config configs/scraper/config.yaml --configExtra  configs/write/config.yaml & 2>/dev/null
	cd web && VITE_READAPI=http://localhost:8087 VITE_WRITEAPI=http://localhost:8088 npm run dev



redis-helm-install:  # Install redis via helm
	./deployment/redis/redis.sh
redis-helm-uninstall:  # Uninstall redis via helm
	helm uninstall redis-stack-server
searchengine-helm-install:  #  Deploy whole microservices to k8s
	helm install searchengine ./deployment/searchengine
searchengine-helm-uninstall:
	helm uninstall searchengine


