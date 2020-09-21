package b_math

import (
	"testing"
)

func TestFloorF(t *testing.T) {
	f := FloorF(2.22)
	if f != 2 {
		t.Fail()
	}
}

func TestCeilF(t *testing.T) {
	f := CeilF(2.22)
	if f != 3 {
		t.Fail()
	}
}

