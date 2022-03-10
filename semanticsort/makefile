BINARY_NAME=semanticsort

# gets the $GOPATH from the shell so we can use it in this makefile
ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

# gets the $GOBIN from the shell so we can use it in this makefile
ifndef $(GOBIN)
    GOBIN=$(shell go env GOBIN)
    export GOBIN
	@echo "Using GOBIN=${GOBIN}"
endif

# puts binary in current folder
build:
	go build

# puts binary in $GOPATH/bin folder
install:
	go install

clean:
	go clean

# cleans/deletes the $GOPATH/pkg/mod folder
cleancache:
	go clean -cache -modcache -i -r

test:
	@echo "Running go test"
	go test

vet:
	@echo "Running go vet"
	go vet
	
lint:
	@echo "Running golangci-lint"
	golangci-lint run --enable-all
