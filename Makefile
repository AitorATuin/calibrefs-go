CWD = $(shell pwd)
PREFIX ?= $(CWD)
GO_VERSION ?= 1.6.2
TAR = $(shell which tar)
GO_URL = https://storage.googleapis.com/golang/go$(GO_VERSION).linux-amd64.tar.gz
GO_TARBALL = /tmp/go_dist.tgz
GO_ROOT = $(CWD)/go_dist
GO_PROJECT = $(CWD)/go_project
GO_PROJECT_NAME = calibrefs
GO = $(GO_ROOT)/go/bin/go

all: go_dist calibrefs

calibrefs_deps:
	@GOPATH=$(PWD) GOROOT=$(GO_ROOT)/go $(GO) get github.com/rminnich/go9p

calibrefs: calibrefs_deps
	@echo "Building calibrefs"
	@GOPATH=$(PWD) GOROOT=$(GO_ROOT)/go $(GO) install calibrefs

go_dist:
	@echo "Getting go binaries"
	@curl $(GO_URL) -o $(GO_TARBALL)
	@mkdir -p $(GO_ROOT)
	@echo "Uncompressing go binaries"
	@$(TAR) zxvf $(GO_TARBALL) -C $(GO_ROOT)

clean_go:
	@$(RM) -rf $(GO_ROOT)

clean_project:
	@$(RM) -rf bin

clean: clean_go  clean_project
