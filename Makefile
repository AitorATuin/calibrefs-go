CWD = $(shell pwd)
PREFIX ?= $(CWD)
GO_VERSION ?= 1.6.2
TAR = $(shell which tar)
GO_URL = https://storage.googleapis.com/golang/go$(GO_VERSION).linux-amd64.tar.gz
GO_TARBALL = /tmp/go_dist.tgz
GO_ROOT = $(CWD)/go_dist
GO_PROJECT = $(CWD)/go_project
GO_PROJECT_NAME = calibrefs
GO_PROJECT_DIR = $(GO_PROJECT)/go/src/github.com/AitorAtuin
GO_PROJECT_BUILD_DIR = $(GO_PROJECT_DIR)
GO_PROJECT_RELPATH = $(shell realpath $(CWD) --relative-to $(GO_PROJECT_BUILD_DIR))
PUSHD = $(shell which pushd)

go_dist:
	@echo "Getting go binaries"
	@curl $(GO_URL) -o $(GO_TARBALL)
	@mkdir -p $(GO_ROOT)
	@echo "Uncompressing go binaries"
	@$(TAR) zxvf $(GO_TARBALL) -C $(GO_ROOT)
	@mkdir -p $(GO_PROJECT_BUILD_DIR)
	@echo "Linking project $(GO_PROJECT_BUILD_DIR) -> $(CWD)"
	@ln -s $(CWD) $(GO_PROJECT_BUILD_DIR)/$(GO_PROJECT_NAME)

clean:
	@$(RM) -rf $(GO_ROOT)
	@$(RM) -rf $(GO_PROJECT)
	@$(RM) -rf $(ARTIFACTS_DIR)
