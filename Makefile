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
# need install Circleci CLI（ref. https://circleci.com/docs/2.0/local-cli/#installation）
ci.l.check:
	${CI} config validate ${CICONF}

ci.l.test:
	${CI} config process ${CICONF} > ${CILOCCONF}
	${CI} build --job test -c ${CILOCCONF}

# For setup gcs on local.
gcs.run:
	docker start --name fake-gcs-server -p 4443:4443 -v ${PWD}/gcs/data:/data fsouza/fake-gcs-server
