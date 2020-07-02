APPNAME=sampleApp
NOW=$(shell date "+%Y-%m-%dT%H:%M:%SZ")
GOCMD=go
GOBUILD=$(GOCMD) build
DOCKERCMD=docker
DOCKERIMAGE=sample_app

build:
	# -X is a option in the go tool link to set value into a variable. (ref. https://golang.org/cmd/link/)
	# -lgflags is a option in the go build to pass arguments to the go tool link. (ref. https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)
	$(GOBUILD) -v -o $(APPNAME) -ldflags="-X 'main.version=v1.0.0' -X 'main.date=$(NOW)'"

run:
	./$(APPNAME)

clean:
	rm $(APPNAME)
	$(GOCMD) clean

build.docker:
	$(DOCKERCMD) build -t $(DOCKERIMAGE) --build-arg VERSION=v1.0.0 .

run.docker:
	$(DOCKERCMD) run $(DOCKERIMAGE)

clean.docker:
	$(DOCKERCMD) image rm -f $(DOCKERIMAGE)
