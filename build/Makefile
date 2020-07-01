APPNAME=sampleApp
NOW=$(shell date "+%Y-%m-%dT%H:%M:%SZ")
GOCMD=go
GOBUILD=$(GOCMD) build

build:
	$(GOBUILD) -v -o $(APPNAME) -ldflags="-X 'main.version=v1.0.0' -X 'main.date=$(NOW)'"

run:
	./$(APPNAME)

clean:
	$(GOCMD) clean
