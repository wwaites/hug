package alg

import (
	"testing"
)

func TestMul(t *testing.T) {
	t.Parallel()

	x := Vector{1,2,3}

	if v := x.Mul(2); !v.Eq(Vector{2,4,6}) {
		t.Errorf("2x -- %s", v)
	}

	if v := x.Mul(0.5); !v.Eq(Vector{0.5,1,1.5}) {
		t.Errorf("x/2 -- %s", v)
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	x := Vector{1,2,3}
	y := Vector{3,2,1}
	z := Vector{4,4,4}

	if v := x.Add(y); !v.Eq(z) {
		t.Errorf("x + y -- %s", v)
	}

	if v := z.Sub(y); !v.Eq(x) {
		t.Errorf("z - y -- %s", v)
	}
}

func TestDot(t *testing.T) {
	t.Parallel()

	x := Vector{1,0,0}
	y := Vector{0,1,0}

	if v := x.Dot(x); v != 1 {
		t.Errorf("x dot x == x -- %f", v)
	}

	if v := x.Dot(y); v != 0 {
		t.Errorf("x dot y == 0 -- %f", v)
	}
}

func TestCross(t *testing.T) {
	t.Parallel()

	x := Vector{1,0,0}
	y := Vector{0,1,0}
	z := Vector{0,0,1}

	n := ZeroV(3)

	if v := x.Cross(x); !v.Eq(n) {
		t.Errorf("x cross x == 0 -- %s", v)
	}

	if v := x.Cross(y); !v.Eq(z) {
		t.Errorf("x cross y == z -- %s", v)
	}

	if v := z.Cross(x); !v.Eq(y) {
		t.Errorf("z cross x == y -- %s", v)
	}

	if v := y.Cross(z); !v.Eq(x) {
		t.Errorf("z cross z == x -- %s", v)
	}
}

func TestNorm(t *testing.T) {
	t.Parallel()

	if v := (Vector{3,4,0}).Length(); v != 5 {
		t.Errorf("||v|| -- %f", v)
	}

	if v := (Vector{0,0,10}).Norm(); !v.Eq(Vector{0,0,1}) {
		t.Errorf("u -- %s", v)
	}
}
