package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"gallows.inf.ed.ac.uk/hug/misc"
	"gallows.inf.ed.ac.uk/hug/proj4"
)

var src_srid int
var dst_srid int

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [opts] x,y[,z]\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.IntVar(&src_srid, "src", 4326, "Source SRID")
	flag.IntVar(&dst_srid, "dst", 27700, "Destination SRID")
	flag.Usage = Usage
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	sproj, err := proj4.InitPlus(fmt.Sprintf("+init=epsg:%d", src_srid))
	if err != nil {
		log.Fatal(err)
	}
	defer sproj.Free()
	dproj, err := proj4.InitPlus(fmt.Sprintf("+init=epsg:%d", dst_srid))
	if err != nil {
		log.Fatal(err)
	}
	defer dproj.Free()

	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "missing argument\n\n")
		Usage()
		os.Exit(255)
	}

	coord, err := misc.ParseCoord(flag.Arg(0))
	if err != nil || len(coord) < 2 || len(coord) > 3 {
		if err != nil {
			log.Print(err)
		}
		Usage()
		os.Exit(255)
	}

	result, err := proj4.Transform(sproj, dproj, coord)
	if err != nil {
		log.Fatal(err)
	}

	outp := make([]string, 0, len(result))
	for _, v := range result {
		outp = append(outp, fmt.Sprintf("%f", v))
	}
	fmt.Print(strings.Join(outp, ","), "\n")
}
