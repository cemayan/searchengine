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



dev: localredis # Start all services


.PHONY: localredis
localredis: # Start a redis-server
	screen -S redis -dm redis-stack-server

localmongogb: # Start a mongodb
	mongod --config /opt/homebrew/etc/mongod.conf --fork
