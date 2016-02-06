.PHONY: all

BIND_DIR := "dist"
LIBKERMIT_MOUNT := -v "$(CURDIR)/$(BIND_DIR):/go/src/github.com/vdemeester/libkermit/$(BIND_DIR)"

GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
LIBKERMIT_DEV_IMAGE := libkermit-dev$(if $(GIT_BRANCH),:$(GIT_BRANCH))
REPONAME := $(shell echo $(REPO) | tr '[:upper:]' '[:lower:]')
INTEGRATION_OPTS := $(if $(MAKE_DOCKER_HOST),-e "DOCKER_HOST=$(MAKE_DOCKER_HOST)", -v "/var/run/docker.sock:/var/run/docker.sock")

DOCKER_RUN_LIBKERMIT := docker run $(if $(CIRCLECI),,--rm) $(INTEGRATION_OPTS) -it $(LIBKERMIT_ENVS) $(LIBKERMIT_MOUNT) "$(LIBKERMIT_DEV_IMAGE)"

default: binary

all: build
	$(DOCKER_RUN_LIBKERMIT) ./hack/make.sh

test: build
	$(DOCKER_RUN_LIBKERMIT) ./hack/make.sh test-unit test-integration

test-integration: build
	$(DOCKER_RUN_LIBKERMIT) ./hack/make.sh test-integration

test-unit: build
	$(DOCKER_RUN_LIBKERMIT) ./hack/make.sh test-unit

validate: build
	$(DOCKER_RUN_LIBKERMIT) ./hack/make.sh validate-gofmt validate-govet validate-golint

validate-gofmt: build
	$(DOCKER_RUN_LIBKERMIT) ./hack/make.sh validate-gofmt

validate-govet: build
	$(DOCKER_RUN_LIBKERMIT) ./hack/make.sh validate-govet

validate-golint: build
	$(DOCKER_RUN_LIBKERMIT) ./hack/make.sh validate-golint

build: dist
	docker build -t "$(LIBKERMIT_DEV_IMAGE)" .

shell: build
	$(DOCKER_RUN_LIBKERMIT) /bin/bash

dist:
	mkdir dist

