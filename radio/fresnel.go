package fresnel

import (
	"math"
	"hug.alg"
)

const C = 299792458.0

func Wavelength(freq float64) float64 {
	return C / freq
}

func Fresnel(x1, x2 alg.Vector, freq float64, n, s int) []alg.Vector {
	link := x1.Sub(x2)
	d := link.Length()
	step := link.Mul(1 / float64(s))
	stepsize := step.Length()
	wavelength := Wavelength(freq)

	ellipsoid := make([]alg.Vector, 0, s+1)
	v := make(alg.Vector, n+1)
	for i, _ := range v {
		v[i] = 0
	}
	ellipsoid = append(ellipsoid, v)
	for i := 1; i <= s; i++ {
		d_1 := float64(i) * stepsize
		d_2 := d - d_1
		v = make(alg.Vector, 1, n+1)
		v[0] = d_1
		for j := 1; j <= n; j++ {
			r := math.Sqrt((float64(j) * wavelength * d_1 * d_2) / (d_1 + d_2))
			v = append(v, r)
		}
		ellipsoid = append(ellipsoid, v)
	}
	return ellipsoid
}
