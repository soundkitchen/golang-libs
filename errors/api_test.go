package errors

import (
	"fmt"
	"testing"
)

// test MultiError with single error.
func TestMultiError1(t *testing.T) {
	m := make(MultiError)
	m["test"] = fmt.Errorf("test")
	actual := m.Error()
	expected := "test: test"
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}

// test MultiError with multiple errors.
func TestMultiError2(t *testing.T) {
	// case multiple errors.
	m := make(MultiError)
	m["test1"] = fmt.Errorf("test1")
	m["test2"] = fmt.Errorf("test2")
	actual := m.Error()
	expected := "test1: test1 (and 1 others)"
	if actual != expected {
		t.Fatalf("expected '%s' but actual '%s'", expected, actual)
	}
}
