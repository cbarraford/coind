
ifndef TARGET
	TARGET=./...
endif

.PHONY: get build test lint install update vet fmt 

install: get
	glide install
	go install . ./cmd/...

update: install

get:
	go get -v ${TARGET}
	go get -u github.com/Masterminds/glide
	go get -v github.com/alecthomas/gometalinter
	gometalinter --install

build:
	go test ${TARGET}

test:
	go test ${TARGET}

vet:
	go vet ${TARGET}

fmt:
	go fmt ${TARGET}

lint: vet fmt
