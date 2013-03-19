package vect

import (
	"fmt"
	"math"
	"strings"
)

// A Vector of floating point numbers
type Vector []float64

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
