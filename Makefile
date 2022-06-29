-include .common.mk

export

all:
	rm -f $(PROJECT)
	$(BUILD_ARGS) go build $(BUILD_FLAGS) ./cmd/$(PROJECT)

codegen: protobuf generate

protobuf: vendor-go protobuf-go-compile

protobuf-go-compile:
	$(foreach SRC, $(PB_GO_SRC), $(shell protoc $(PB_OPTIONS) $(PB_GO_INCLUDE) $(PB_GO_COMPILE) `ls $(SRC)/*.proto`))

generate:
	go generate ./...

vendor-go:
	go mod tidy
	go mod vendor

test:
	go test $(TESTS_OPTS) $(TESTS)

test-all: test-env
	$(DOCKER_COMPOSE) up make

test-env:
	$(DOCKER_COMPOSE) create make

test-clean:
	$(DOCKER_COMPOSE) down --remove-orphans

tool-chain:
	go install \
		github.com/kisielk/errcheck \
		github.com/maxbrunsfeld/counterfeiter/v6 \
		github.com/rleszilm/genms-version \
		golang.org/x/lint/golint \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		honnef.co/go/tools/cmd/staticcheck

lint:
	golint `go list ./...`
	staticcheck `go list ./... | grep -v github.com/rleszilm/genms-datalayer/tools`
	errcheck -ignoretests -ignoregenerated -asserts -exclude .errcheck ./...
