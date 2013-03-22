package alg

import (
	"testing"
)

func TestDot(t *testing.T) {
	t.Parallel()

	x := Vector{1,0,0}
	y := Vector{0,1,0}

	if v := x.Dot(x); v != 1 {
		t.Errorf("x dot x == x -- %.16f", v)
	}

	if v := x.Dot(y); v != 0 {
		t.Errorf("x dot y == 0 -- %.16f", v)
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
