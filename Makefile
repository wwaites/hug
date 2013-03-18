GCCGO  ?=gccgo
CFLAGS ?=-O3 -g -Wall -Werror
OBJS   =\
	hutil.o \
	cproj.o \
	jproj.o \
	server.o
PROG=geosrv

all: ${PROG}

clean:
	rm -f *.o *~
	rm -f ${PROG}

%.o: %.go
	${GCCGO} ${CFLAGS} -c -o $@ $<

${PROG}: ${OBJS}
	${GCCGO} -o $@ $^ -lgo -lproj
