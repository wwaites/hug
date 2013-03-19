package main

import (
	"fmt"
	"fresnel"
	"geo"
)

func main() {
	x1 := geo.Cartesian(geo.LonLat{-5, 58}, geo.R)
	x2 := geo.Cartesian(geo.LonLat{-4.5, 58}, geo.R)
	zone := fresnel.Fresnel(x1, x2, 5.8e9, 3, 100)
	for _, x := range zone {
		fmt.Printf("%s\n", x)
		fmt.Printf("%f %f %f %f\n", x[0], x[1], x[2], x[3])
	}
}