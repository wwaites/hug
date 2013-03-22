GCCGO        ?=gccgo
GCCGOCFLAGS  ?=-O3 -g -Wall -Werror -I/usr/lib/go/pkg/gccgo -lproj

all: build

build:
	go build -compiler=${GCCGO} -gccgoflags="${GCCGOFLAGS}" gallows.inf.ed.ac.uk/hug/...

install:
	go install -compiler=${GCCGO} -gccgoflags="${GCCGOFLAGS}" gallows.inf.ed.ac.uk/hug/...

test:
	go test -compiler=${GCCGO} -gccgoflags="${GCCGOFLAGS}" gallows.inf.ed.ac.uk/hug/...
