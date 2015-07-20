NAME = mock-ass
MOCK_ASS_DATA_DIR = $$(pwd)/data

env:
	@echo \#!/bin/bash > env.sh
	@echo export GOPATH=$$(pwd):$$(pwd)/vendor >> env.sh
	@echo export MOCK_ASS_DATA_DIR=${MOCK_ASS_DATA_DIR} >> env.sh
	@echo exec \$$@ >> env.sh
	chmod a+x env.sh

all: build

build:
	gb build

run: env
	./env.sh bin/$(NAME)

build-run: build run

run-dev: env
	./env.sh go run ./src/mock-ass/main.go -color
