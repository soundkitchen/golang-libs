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
	ReAny                = regexp.MustCompile("^[!-~]+$")
)

//
func TestPickupDigits(t *testing.T) {
	expected := "7980019456687858161603743195850682668430"
	actual := Pickup(40, FDigits)
	if !ReDigits.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupLowerCase(t *testing.T) {
	expected := "jjpremtftszjanerefzckbkqltmtontynaejijsc"
	actual := Pickup(40, FLowerCase)
	if !ReLowerCase.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupUpperCase(t *testing.T) {
	expected := "NWNGUSMOFBBJFGSWKMOGOBOEKPOEVAZNXECKXGRO"
	actual := Pickup(40, FUpperCase)
	if !ReUpperCase.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupPunctuation(t *testing.T) {
	expected := "|&=+~#/~]_%^<+|;%@|~.__^^{=}'>|:\\|^;;>{?"
	actual := Pickup(40, FPunctuation)
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupDigitsAndLowerCase(t *testing.T) {
	expected := "5s1g2kiiu3m0is36m1i556y4z7oomnrnnw5mdj54"
	actual := Pickup(40, FDigits|FLowerCase)
	if !ReDigitsAndLowerCase.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupDigitsAndUpperCase(t *testing.T) {
	expected := "32U0SGHL891G1BG4QA52UJ1373UM00YADOISGY09"
	actual := Pickup(40, FDigits|FUpperCase)
	if !ReDigitsAndUpperCase.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupDigitsAndPunctuation(t *testing.T) {
	expected := "6#|;9<_7=969_`/:(2?(11+2|40_266#|:718@2|"
	actual := Pickup(40, FDigits|FPunctuation)
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupLowerCaseAndPunctuation(t *testing.T) {
	expected := "s\\|ain&sq/nzazi=\\a@{pv+ixf<}~d~y@o\\=s&qr"
	actual := Pickup(40, FLowerCase|FPunctuation)
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupUpperCaseAndPunctuation(t *testing.T) {
	expected := "NB~L>>([QZQF)N_>KL>$S`HV{ZE']X!Z|N:YSQNT"
	actual := Pickup(40, FUpperCase|FPunctuation)
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupLetters(t *testing.T) {
	expected := "gAcOVYApIXkamDZCfSQxlfAAZDudIGVRpwrXTVnD"
	actual := Pickup(40, FLetters)
	if !ReLetters.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

//
func TestPickupAny(t *testing.T) {
	expected := "0\\*Dl3qLEG8xb]*76BDRPyYF1-~tKjQ)DFge{cf8"
	actual := Pickup(40, FAny)
	if !ReAny.MatchString(actual) {
		t.Fatalf("invalid value: %s", actual)
	}
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}
