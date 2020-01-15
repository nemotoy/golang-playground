GOCMD=go
GOTEST=$(GOCMD) test
GOLINT=golangci-lint

all: lint test

test:
	$(GOTEST) -v ./...

test.race:
	$(GOTEST) -v -race ./...

ci.test:
	$(GOTEST) -race -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	$(GOLINT) run -v ./...

# For local
ci.l.check:
	circleci config validate .circleci/config.yml

ci.l.test:
	circleci config process .circleci/config.yml > ci_local.yml
	circleci build --job test -c ci_local.yml
