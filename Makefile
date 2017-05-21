VERSION=0.0.1
#BUILD=`git rev-parse HEAD`

#LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
LDFLAGS=-ldflags "-s -w"

build: deps
	go build ${LDFLAGS}

check:
	go fmt -x
	go vet -x

test: check
	go test

deps:
	go get -d -v -u -f

install: deps
	go install ${LDFLAGS}

clean:
	go clean -i -x

.PHONY: clean install deps test check build

# vim: set noexpandtab shiftwidth=8 softtabstop=0:
