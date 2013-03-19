GCCGO  ?=gccgo
CFLAGS ?=-O3 -g -Wall -Werror
OBJS   =\
	webx.o \
	cproj.o \
	jproj.o \
	vect.o \
	geo.o \
	geosrv.o 
PROG=geosrv

all: ${PROG}

clean:
	rm -f *.o *~
	rm -f ${PROG}

%.o: %.go
	${GCCGO} ${CFLAGS} -c -o $@ $<

${PROG}: ${OBJS}
	${GCCGO} -o $@ $^ -lgo -lproj

gt: gt.o vect.o geo.o
	${GCCGO} -o $@ $^ -lgo
	./gt > nn.dat
	gnuplot nn.plt

fr: vect.o geo.o fresnel.o fr.o
	${GCCGO} -o $@ $^ -lgo
	./fr > fr.dat
	gnuplot fr.plt
