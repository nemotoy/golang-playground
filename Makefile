GOCMD=go
GOTEST=$(GOCMD) test
GOLINT=golangci-lint
CI=circleci
CICONF=.circleci/config.yml
CILOCCONF=ci_local.yml

all: lint test

test:
	$(GOTEST) -v ./...

test.race:
	$(GOTEST) -v -race ./...

ci.test:
	$(GOTEST) -race -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	$(GOLINT) run -v ./... -c .golangci.yaml

# For local
ci.l.check:
	${CI} config validate ${CICONF}

ci.l.test:
	${CI} config process ${CICONF} > ${CILOCCONF}
	${CI} build --job test -c ${CILOCCONF}
