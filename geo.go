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
func ChordHeight(p1, p2 LonLat, r float64, n int) (ch chan vect.Vector) {
	// transform to cartesian coordinates
	x1 := Cartesian(p1, r)
	x2 := Cartesian(p2, r)

	// find the step size
	d := x2.Sub(x1)
	step := d.Mul(1 / float64(n))
	stepsize := step.Length()

	// implementation detail -- this is just AdjustAlt with altitude set to zero!
	pts := make(chan vect.Vector)
	go func() {
		for i := 0 ; i <= n; i++ {
			pts <- vect.Vector{float64(i) * stepsize, 0}
		}
		close(pts)
	}()
	
	return AdjustAlt(p1, p2, pts, r)
}

// Given two points, as with ChordHeight, and a sequence of data points representing 
// the terrain height between those two points, adjust the data to correct for the curvature
// of the earth.
func AdjustAlt(p1, p2 LonLat, pts chan vect.Vector, r float64) (ch chan vect.Vector) {
	ch = make(chan vect.Vector)

	go func () {
		// transform to cartesian coordinates
		x1 := Cartesian(p1, r)
		x2 := Cartesian(p2, r)
	
		// unit vector in the direction between the two points
		ud := x2.Sub(x1).Norm()
		// unit vector in the vertical direction
		uh := ud.Cross(x1).Cross(ud).Norm()

		for pt := range pts {
			x := x1.Add(ud.Mul(pt[0]))
			h := x.Norm().Mul(pt[1] + r)
			v := vect.Vector{
				// component in the direction of the line
				h.Dot(ud) - x1.Dot(ud),
				// component perpendicular to the line
				h.Dot(uh) - x.Dot(uh),
			}
			ch <- v
		}

		close(ch)
	}()

	return ch
}