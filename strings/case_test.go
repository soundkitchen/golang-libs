package strings

import (
	"fmt"
	"testing"
)

func TestToCamelCase(t *testing.T) {
	cases := []struct {
		Initial  string
		Expected string
	}{
		{Initial: "spam", Expected: "Spam"},
		{Initial: "who_says_ni", Expected: "WhoSaysNi"},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			actual := ToCamelCase(c.Initial)
			if actual != c.Expected {
				t.Fatalf("expected %v but actual %v", c.Expected, actual)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	cases := []struct {
		Initial  string
		Expected string
	}{
		{Initial: "Spam", Expected: "spam"},
		{Initial: "EGG", Expected: "egg"},
		{Initial: "WhoSaysNi", Expected: "who_says_ni"},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			actual := ToSnakeCase(c.Initial)
			if actual != c.Expected {
				t.Fatalf("expected %v but actual %v", c.Expected, actual)
			}
		})
	}
}
