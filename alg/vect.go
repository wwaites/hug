package alg

import (
	"fmt"
	"math"
	"strings"
)

// A Vector of floating point numbers
type Vector []float64

// Return a zero vector of dimension n
func ZeroV(n int) Vector {
	v := make(Vector, n)
	for i := 0; i < n; i++ {
		v[i] = 0
	}
	return v

}

// Test if the two vectors are equal
func (v1 Vector) Eq(v2 Vector) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i, x := range v1 {
		if x != v2[i] {
			return false
		}
	}
	return true
}

// Scalar multiplication of a vector
func (v Vector) Mul(a float64) Vector {
	v2 := make(Vector, 0, len(v))
	for _, x := range v {
		v2 = append(v2, x * a)
	}
	return v2
}

// Vector addition
func (v1 Vector) Add(v2 Vector) Vector {
	v := make(Vector, 0, len(v1))
	for i, x := range v1 {
		v = append(v, x + v2[i])
	}
	return v
}

// Vector subtraction (implemented separately for efficiency)
func (v1 Vector) Sub(v2 Vector) Vector {
	v := make(Vector, 0, len(v1))
	for i, x := range v1 {
		v = append(v, x - v2[i])
	}
	return v
}

// Scalar or dot product
func (v1 Vector) Dot(v2 Vector) float64 {
	var product float64
	for i, x := range v1 {
		product += x * v2[i]
	}
	return product
}

// Vector product XXX only for 3 dimensional vectors!
func (a Vector) Cross(b Vector) Vector {
	c := make(Vector, 3)
	c[0] = a[1] * b[2] - a[2] * b[1]
	c[1] = a[2] * b[0] - a[0] * b[2]
	c[2] = a[0] * b[1] - a[1] * b[0]
	return c
}

// The length of the vector
func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

// Normalise the vector to unit length
func (v Vector) Norm() Vector {
	return v.Mul(1 / v.Length())
}

// A nice, pretty, string representation
func (v Vector) String() string {
	s := make([]string, 0, len(v))
	for _, v := range v {
		s = append(s, fmt.Sprintf("%.06f", v))
	}
	return "[" + strings.Join(s, ", ") + "]"
}

type Matrix []Vector
func (m Matrix) Transpose() (r Matrix) {
	r = make(Matrix, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		r[i] = make(Vector, len(m))
		for j := 0; j < len(r[i]); j++ {
			r[i][j] = m[j][i]
		}
	}
	return r
}

func (m1 Matrix) Mul(m2 Matrix) (r Matrix) {
	m2t := m2.Transpose()
	r = make(Matrix, len(m2))
	for i := 0; i < len(m1); i++ {
		r[i] = make(Vector, len(m2t))
		for j := 0; j < len(m2t); j++ {
			r[i][j] = m1[i].Dot(m2t[j])
		}
	}
	return r
}

func (m Matrix) String() string {
	vs := make([]string, 0, len(m))
	for _, v := range m {
		vs = append(vs, v.String())
	}
	return "[" + strings.Join(vs, ", ")  + "]"
}
