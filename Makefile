GCCGO       ?=gccgo
GCCGOFLAGS  ?=-O3 -pthread -fPIC -g -Wall -Werror -I/usr/lib/go/pkg/gccgo -lproj
GOROOT      ?=/usr/lib/go

all: build

build:
	@go build -x -v -compiler=${GCCGO} -gccgoflags="${GCCGOFLAGS}" gallows.inf.ed.ac.uk/hug/...

install:
	mkdir -p ${GOROOT}/pkg/gccgo
	mkdir -p ${GOROOT}/bin
	@go install -v -compiler=${GCCGO} -gccgoflags="${GCCGOFLAGS}" gallows.inf.ed.ac.uk/hug/...

test:
	@go test -v -compiler=${GCCGO} -gccgoflags="${GCCGOFLAGS}" gallows.inf.ed.ac.uk/hug/...

rebuild:
	@find . -name \*.go -exec touch '{}' ';'
	@${MAKE} build
