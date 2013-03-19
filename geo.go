package geo

import (
	"fmt"
	"math"
	"vect"
)

// Radius of the earth in meters
const R = 6378100.0

type LonLat vect.Vector
func (l LonLat)Lon() float64 {
	return l[0]
}
func (l LonLat)Lat() float64 {
	return l[1]
}
func (l LonLat)String() string {
	return fmt.Sprintf("(%f, %f)", l[0], l[1])
}

// Transform from spherical coordinates in the geographic convention
// to cartesian coordinates with the x axis pointing towards the prime
// meridian and the z axis pointing to the north pole
func Cartesian(p LonLat, r float64) vect.Vector {
	x := make(vect.Vector, 3)
	// convert to radians
	lon := p.Lon() * math.Pi / 180
	lat := p.Lat() * math.Pi / 180
	// convert to more "standard" spherical coordinates
	theta := math.Pi / 2 - lat  // inclination from z axis
	phi   := lon                // azimuth from x axis
	// transform
	x[0] = r * math.Sin(theta) * math.Cos(phi)
	x[1] = r * math.Sin(theta) * math.Sin(phi)
	x[2] = r * math.Cos(theta)
	return x
}

// Given two points on a sphere, expressed using the usual geographical convention
// for latitude and longitude, and the radius of the sphere, calculate the distance between
// the chord or direct line between the two points and the surface with a resolution of n.
// The return value is a list of (d,h) pairs where d is the distance along the chord starting
// at p1 and h is the difference between r and the distance from the centre of the sphere
// to that point.
func ChordHeight(p1, p2 LonLat, r float64, n int) []vect.Vector {
	// transform to cartesian coordinates
	x1 := Cartesian(p1, r)
	x2 := Cartesian(p2, r)
	
	d := x2.Sub(x1)
	step := d.Mul(1 / float64(n))
	stepsize := step.Length()

	curve := make([]vect.Vector, 0, n + 1)
	curve = append(curve, vect.Vector{0,0})

	// Starting at x1, incrementally move towards x2
	// recording the difference between r and x, i.e. the distance from
	// the chord to the circle
	for i, x := 1, x1; i < n+1; i++ {
		x = x.Add(step)
		pt := vect.Vector{float64(i) * stepsize, r - x.Length()}
		curve = append(curve, pt)
	}
	return curve
}