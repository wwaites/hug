package main

import (
	"fmt"
	"geo"
//	"encoding/json"
)

func main() {
/*
	ll := geo.LonLat{0,0}
	fmt.Printf("%s -> %s\n", ll, geo.Cartesian(ll, 1))
	ll = geo.LonLat{0,90}
	fmt.Printf("%s -> %s\n", ll, geo.Cartesian(ll, 1))
	ll = geo.LonLat{90,0}
	fmt.Printf("%s -> %s\n", ll, geo.Cartesian(ll, 1))
	ll = geo.LonLat{90,90}
	fmt.Printf("%s -> %s\n", ll, geo.Cartesian(ll, 1))

	return
*/
	c := geo.ChordHeight(geo.LonLat{0,0}, geo.LonLat{0, 90}, 1.0, 10)
//	fmt.Print(c, "\n")

	c = geo.ChordHeight(geo.LonLat{90,0}, geo.LonLat{0, 90}, geo.R, 100)
//	c = geo.ChordHeight(geo.LonLat{-5.9,0}, geo.LonLat{-6, 0}, geo.R, 100)
	for _, v := range c {
		fmt.Printf("%f %f\n", v[0], v[1])
	}
//	fmt.Print(c)
//	fmt.Print(c, "\n")

//	b, err := json.Marshal(c)
//	if err != nil {
//		fmt.Print(err)
//	}
//	fmt.Printf("%s", b)
}
