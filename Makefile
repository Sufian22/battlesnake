export APP = battlesnake
export CONFIG_FILE ?= "config.json"

.PHONY: all
all: help

.PHONY: help
help:
	@echo "make clean - clean test cache, build files"
	@echo "make build - build binary"
	@echo "make test - run go test including race detection"
	@echo "make build-docker - creates a docker image"
	@echo "make run-docker - runs a docker image"

.PHONY: clean
clean:
	@rm -rf ./build
	@go clean ${CLEAN_OPTIONS}

.PHONY: build
build: clean
	@go build -tags netgo -a -v -o ./build/${APP} cmd/battlesnake/main.go
	@chmod +x ./build/*

.PHONY: test
test: 
	@go test ./... -cover

.PHONY: run
run: 
	@go run cmd/battlesnake/main.go -config ${CONFIG_FILE}

.PHONY: build-docker
build-docker:
	@docker build -t ${APP} .

.PHONY: run-docker
run-docker: build-docker
	@docker run --rm -p 3000:3000 -t ${APP} 