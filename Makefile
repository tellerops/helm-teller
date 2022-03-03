# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTOOL=$(GOCMD) tool
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD)fmt

BINARY_NAME=helm-teller

TEST_EXEC_CMD=$(GOTEST) -coverprofile=cover.out -short -cover -failfast ./... 

build: build
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(TEST_EXEC_CMD)
		
test-html:
	$(TEST_EXEC_CMD)
	$(GOTOOL) cover -html=cover.out

checks: test lint fmt

lint:
	golangci-lint run

fmt: 
	@res=$$($(GOFMT) -d -e -s $$(find . -type d \( -path ./src/vendor \) -prune -o -name '*.go' -print)); \
	if [ -n "$${res}" ]; then \
		echo checking gofmt fail... ; \
		echo "$${res}"; \
		exit 1; \
	fi

build: 
	goreleaser release --snapshot --rm-dist