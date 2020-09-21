package b_math

import (
	"testing"
)

func TestMaxI(t *testing.T) {
	v := MaxI(1, 2, 3, 4, 7, 8, 0)
	if v != 8 {
		t.Fail()
	}
}

func TestMinI(t *testing.T) {
	v := MinI(1, 2, 3, 4, 7, 8, 0)
	if v != 0 {
		t.Fail()
	}
}

func TestMaxF(t *testing.T) {
	v := MaxF(1.0, 2, 3, 4, 7, 8.1, 0)
	if v != 8.1 {
		t.Fail()
	}
}

func TestMinF(t *testing.T) {
	v := MinF(1.0, 2, 3, 4, 7, 8.1, 0)
	if v != 0 {
		t.Fail()
	}
}
