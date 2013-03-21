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
UTILS=tgu_convert tgu_fresnel tgu_circle tgu_addseq tgu_curve

all: ${PROG} ${UTILS}

clean:
	rm -f *.o *~
	rm -f ${PROG} ${UTILS}

%.o: %.go
	${GCCGO} ${CFLAGS} -c -o $@ $<

${PROG}: ${OBJS}
	${GCCGO} -o $@ $^ -lgo -lproj

tgu_convert: vect.o cproj.o tgu_convert.o
	${GCCGO} -o $@ $^ -lgo -lproj

tgu_fresnel: vect.o geo.o fresnel.o cproj.o tgu_fresnel.o
	${GCCGO} -o $@ $^ -lgo -lproj

tgu_circle: vect.o geo.o cproj.o tgu_circle.o
	${GCCGO} -o $@ $^ -lgo -lproj

tgu_addseq: vect.o tgutil.o tgu_addseq.o
	${GCCGO} -o $@ $^ -lgo

tgu_curve: vect.o geo.o cproj.o tgutil.o tgu_curve.o
	${GCCGO} -o $@ $^ -lgo -lproj
