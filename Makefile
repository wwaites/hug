GCCGO       ?=gccgo
GCCGOFLAGS  ?=-O3 -g -Wall -Werror -I/usr/lib/go/pkg/gccgo -lproj

all: build

build:
	@go build -v -compiler=${GCCGO} -gccgoflags="${GCCGOFLAGS}" gallows.inf.ed.ac.uk/hug/...

install:
	@go install -v -compiler=${GCCGO} -gccgoflags="${GCCGOFLAGS}" gallows.inf.ed.ac.uk/hug/...

test:
	@go test -v -compiler=${GCCGO} -gccgoflags="${GCCGOFLAGS}" gallows.inf.ed.ac.uk/hug/...

rebuild:
	@find . -name \*.go -exec touch '{}' ';'
	@${MAKE} build
