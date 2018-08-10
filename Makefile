
ifndef TARGET
	TARGET=$(shell glide novendor)
endif

.PHONY: get build test lint install update vet fmt 

install: get
	glide install
	go install . ./cmd/...

update: install

get:
	go get -v ${TARGET}
	go get -u github.com/Masterminds/glide
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter --install

build:
	go build ${TARGET}

test:
	go test ${TARGET}

lint: vet fmt
	gometalinter.v2 -j 4 --disable-all --enable=gofmt --enable=golint --enable=vet --enable=gosimple --enable=unconvert --deadline=10m ${TARGET} 2>&1 | grep -v 'ALL_CAPS\|OP_' 2>&1
