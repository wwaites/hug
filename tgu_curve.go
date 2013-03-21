package main

import (
	"cproj"
	"flag"
	"fmt"
	"geo"
	"log"
	"os"
	"tgutil"
	"vect"
)

var srid int
var r float64
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [opts] x1,y1 x2,y2 filename.dat\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.IntVar(&srid, "srid", 4326, "SRID for the given points")
	flag.Float64Var(&r, "r", geo.R, "Radius of sphere for geographical coordinates (m)")
	flag.Usage = Usage
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	if len(flag.Args()) != 3 {
		Usage()
		os.Exit(255)
	}

	p1, err := vect.ParseCoord(flag.Arg(0))
	if err != nil || len(p1) < 2 {
		if err != nil {
			log.Print(err)
		}
		Usage()
		os.Exit(255)
	}

	p2, err := vect.ParseCoord(flag.Arg(1))
	if err != nil || len(p2) < 2 {
		if err != nil {
			log.Print(err)
		}
		Usage()
		os.Exit(255)
	}

	var ll1, ll2 geo.LonLat
	if srid == 4326 {
		ll1, ll2 = geo.LonLat(p1), geo.LonLat(p2)
	} else {
		proj, err := cproj.InitPlus(fmt.Sprintf("+init=epsg:%d", srid))
		if err != nil {
			log.Fatal(err)
		}
		defer proj.Free()

		wgs84, err := cproj.InitPlus("+init=epsg:4326")
		if err != nil {
			log.Fatal(err)
		}
		defer wgs84.Free()

		ll1, err = cproj.Transform(proj, wgs84, p1)
		if err != nil {
			log.Fatal(err)
		}
		ll2, err = cproj.Transform(proj, wgs84, p2)
		if err != nil {
			log.Fatal(err)
		}
	}

	fp, err := os.Open(flag.Arg(2))
	if err != nil {
		log.Fatal(err)
	}

	n := tgutil.Numbers(fp)

	for v := range geo.AdjustAlt(ll1, ll2, n, r) {
		fmt.Printf("%f %f\n", v[0], v[1])
	}
}