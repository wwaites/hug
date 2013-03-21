package main

import (
	"cproj"
	"flag"
	"fmt"
	"fresnel"
	"geo"
	"log"
	"math"
	"os"
	"vect"
)

var s, srid int
var r, freq float64
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [opts] x1,y1[,z1] x2,y2[,z2]\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.IntVar(&srid, "srid", 4326, "SRID for the given points")
	flag.Float64Var(&r, "r", geo.R, "Radius of sphere for geographical coordinates (m)")
	flag.Float64Var(&freq, "freq", 5.5e9, "Frequency for calculation (Hz)")
	flag.IntVar(&s, "s", 10, "Step distance between output points (m)")
	flag.Usage = Usage
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	if len(flag.Args()) != 2 {
		Usage()
		os.Exit(255)
	}

	p1, err := vect.ParseCoord(flag.Arg(0))
	if err != nil || len(p1) < 2 || len(p1) > 3 {
		if err != nil {
			log.Print(err)
		}
		Usage()
		os.Exit(255)
	}

	p2, err := vect.ParseCoord(flag.Arg(1))
	if err != nil || len(p2) < 2 || len(p2) > 3 {
		if err != nil {
			log.Print(err)
		}
		Usage()
		os.Exit(255)
	}

	var ll1, ll2 vect.Vector
	if srid == 4326 {
		ll1, ll2 = p1, p2
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

	var h1, h2 float64
	if len(p1) == 2 {
		h1 = 0
	} else {
		h1 = p1[2]
	}
	if len(p2) == 2 {
		h2 = 0
	} else {
		h2 = p2[2]
	}
	x1 := geo.Cartesian(geo.LonLat(ll1), r + h1)
	x2 := geo.Cartesian(geo.LonLat(ll2), r + h2)

	// find the tangent vector so we can find the elevation of the link...
	tan := geo.Cartesian(geo.LonLat(ll1), r).Sub(
		geo.Cartesian(geo.LonLat(ll2), r))
	link := x1.Sub(x2)
	costheta := tan.Dot(link) / tan.Length() / link.Length()
	theta := math.Acos(costheta)
//	slope := math.Tan(theta)

	R := vect.Matrix{
		{ math.Cos(theta), -1 * math.Sin(theta) },
		{ math.Sin(theta), math.Cos(theta) },
	}

	if h1 > h2 {
		R = R.Transpose()
	}

	steps := int(link.Length()) / s + 1
	ellipsoid := fresnel.Fresnel(x1, x2, freq, 1, steps)
	for _, v := range ellipsoid {
		// correct for the altitude at each of the points...
		// elevation of the link itself
		//le := v[0] * slope + h1
		// and rotate the point itself
		x0 := vect.Matrix{
			{ v[0] },
			{ h1 },
		}
		x1 := vect.Matrix {
			{ v[0] },
			{ h1 - 1 * v[1] },
		}
		x2 := vect.Matrix {
			{ v[0] },
			{ h1 + v[1] },
		}

		x0r := R.Mul(x0)
		fmt.Printf("%f %f\n", x0r[0][0], x0r[1][0])

		x1r := R.Mul(x1)
		fmt.Printf("%f %f\n", x1r[0][0], x1r[1][0])

		x2r := R.Mul(x2)
		fmt.Printf("%f %f\n", x2r[0][0], x2r[1][0])
	}
}
