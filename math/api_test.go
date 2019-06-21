package math

import (
	"testing"
)

var fixtures = []interface{}{-10, 0, 5.5, 100.5}

//
func TestMax(t *testing.T) {
	a := Max(fixtures...)
	if a != 100.5 {
		t.Fatalf("expected 100.5 but actual %f", a)
	}
}

//
func TestMin(t *testing.T) {
	a := Min(fixtures...)
	if a != -10 {
		t.Fatalf("expected -10 but actual %f", a)
	}
}

//
func TestSum(t *testing.T) {
	a := Sum(fixtures...)
	if a != 96 {
		t.Fatalf("expected 96 but actual %f", a)
	}
}

//
func TestAverage(t *testing.T) {
	a := Average(fixtures...)
	if a != 24 {
		t.Fatalf("expected 24 but actual %f", a)
	}
}
