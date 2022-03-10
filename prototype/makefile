BINARY_NAME=prototype

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
	go build -o ${BINARY_NAME} main.go

# puts binary in $GOPATH/bin folder
install:
	go install

# puts binary in $GOPATH/bin/${BINARY_NAME} folder
deploy:
	@echo "Making folder ${GOBIN}/modulemonkey/prototype/${BINARY_NAME}"
	mkdir -p ${GOBIN}/modulemonkey/prototype/${BINARY_NAME}
	@echo "Copying ${BINARY_NAME} to ${GOBIN}/modulemonkey/prototype/${BINARY_NAME}"
	cp ${BINARY_NAME} ${GOBIN}/modulemonkey/prototype/${BINARY_NAME}

# builds a docker image and installs to the local docker server
docker:
	@echo "Running docker build with BIN_PATH=${GOBIN}/${BINARY_NAME} and BIN_FILE=${BINARY_NAME}"
	docker build --build-arg BIN_PATH=${GOBIN}/${BINARY_NAME}/ --build-arg BIN_FILE=${BINARY_NAME} .

clean:
	@echo "Running go clean for BINARY_NAME = ${BINARY_NAME}"
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

