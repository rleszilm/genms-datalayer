PROJECT ?= genms-datalayer
GITHUB_ACCOUNT ?= github.com/rleszilm
GITHUB_REPO ?= $(GITHUB_ACCOUNT)/$(PROJECT)

export GOPRIVATE=$(GITHUB_ACCOUNT)

ARANGODB_VERSION ?= 3.9.1
GO_VERSION ?= 1.18.2

BUILD_BRANCH_OVERRIDE ?= PR-${TRAVIS_PULL_REQUEST}
BUILD_RELEASE ?= $(shell genms-version minor -fis -B $(BUILD_BRANCH_OVERRIDE))
BUILD_REVISION ?= $(shell genms-version minor -firs -B $(BUILD_BRANCH_OVERRIDE))

BUILD_ARGS ?= CGO_ENABLED=0
BUILD_FLAGS ?= -a -ldflags "-s \
	-X $(GITHUB_REPO)/build.release=$(BUILD_RELEASE) \
	-X $(GITHUB_REPO)/build.revision=$(BUILD_REVISION)"

PB_OPTIONS := --experimental_allow_proto3_optional
PB_GO_INCLUDE := -I pkg
PB_GO_COMPILE := --go_opt=paths=source_relative \
	--go_out=pkg
PB_GO_SRC := pkg/annotations pkg/annotations/bson

TEST_OPTS ?= -race -coverprofile=cover.out
TEST_EXCLUDE_PACKAGES = $(GITHUB_REPO)/tools

TESTS ?= $(TEST_OPTS) $(shell go list ./... | grep -v $(TEST_EXCLUDE_PACKAGES))
