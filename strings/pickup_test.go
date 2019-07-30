package strings

import (
	"regexp"
	"testing"
)

var (
	ReDigits             = regexp.MustCompile("^[0-9]+$")
	ReLowerCase          = regexp.MustCompile("^[a-z]+$")
	ReUpperCase          = regexp.MustCompile("^[A-Z]+$")
	ReDigitsAndLowerCase = regexp.MustCompile("^[0-9a-z]+$")
	ReDigitsAndUpperCase = regexp.MustCompile("^[0-9A-Z]+$")
	ReLetters            = regexp.MustCompile("^[a-zA-Z]+$")
	ReAny                = regexp.MustCompile("^[0-9a-zA-Z]+$")
)

//
func TestPickupDigits(t *testing.T) {
	expected := "79800194"
	actual := Pickup(8, FDigits)
	if !ReDigits.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupLowerCase(t *testing.T) {
	expected := "tccahkfa"
	actual := Pickup(8, FLowerCase)
	if !ReLowerCase.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupUpperCase(t *testing.T) {
	expected := "PSFCOFRW"
	actual := Pickup(8, FUpperCase)
	if !ReUpperCase.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupDigitsAndLowerCase(t *testing.T) {
	expected := "pdj58vge"
	actual := Pickup(8, FDigits|FLowerCase)
	if !ReDigitsAndLowerCase.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupDigitsAndUpperCase(t *testing.T) {
	expected := "8EQE8YR0"
	actual := Pickup(8, FDigits|FUpperCase)
	if !ReDigitsAndUpperCase.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupLetters(t *testing.T) {
	expected := "JJPREmtf"
	actual := Pickup(8, FLetters)
	if !ReLetters.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupAny(t *testing.T) {
	expected := "3s17An23"
	actual := Pickup(8, FAny)
	if !ReAny.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}
