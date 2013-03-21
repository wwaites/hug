GCCGO   ?=gccgo
CFLAGS  ?=-O3 -g -Wall -Werror
INSTALL ?=/usr/bin/install
PREFIX  ?=/usr/local
DAEMON  =tgud
UTILS   =tgu_convert tgu_fresnel tgu_circle tgu_addseq tgu_curve
SCRIPTS =tgu_rflos

all: ${DAEMON} ${UTILS}

clean:
	rm -f *.o *~
	rm -f ${DAEMON} ${UTILS}

install: all
	for prog in ${UTILS}; do \
		${INSTALL} -c -m 755 $$prog ${PREFIX}/bin/$$prog; \
	done
	for prog in ${SCRIPTS}; do \
		${INSTALL} -c -m 755 $$prog ${PREFIX}/bin/$$prog; \
	done
	${INSTALL} -c -m 755 ${DAEMON} ${PREFIX}/sbin/${DAEMON}

%.o: %.go
	${GCCGO} ${CFLAGS} -c -o $@ $<

tgud: webx.o cproj.o jproj.o vect.o geo.o geosrv.o
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
