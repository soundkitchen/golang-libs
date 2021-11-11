package math

import (
	"testing"
)

var positiveNumbers = []interface{}{1, 2.5, 3.14159, 10}
var negativeNumbers = []interface{}{-10, -5.2, -3.14159, -1}
var fixtures = []interface{}{-10, 0, 5.5, 100.5}

//
func TestMax(t *testing.T) {
	t.Run("with positive numbers.", func(t *testing.T) {
		actual := Max(positiveNumbers...)
		expected := 10.0
		if actual != expected {
			t.Errorf("expected %f but actual %f", expected, actual)
		}
	})
	t.Run("with negative numbers.", func(t *testing.T) {
		actual := Max(negativeNumbers...)
		expected := -1.0
		if actual != expected {
			t.Errorf("expected %f but actual %f", expected, actual)
		}
	})
}

//
func TestMin(t *testing.T) {
	t.Run("with positive numbers.", func(t *testing.T) {
		actual := Min(positiveNumbers...)
		expected := 1.0
		if actual != expected {
			t.Errorf("expected %f but actual %f", expected, actual)
		}
	})
	t.Run("with negative numbers.", func(t *testing.T) {
		actual := Min(negativeNumbers...)
		expected := -10.0
		if actual != expected {
			t.Errorf("expected %f but actual %f", expected, actual)
		}
	})
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
